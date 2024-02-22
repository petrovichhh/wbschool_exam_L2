package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

var LastID = 0
var LastIdMutex = sync.Mutex{}

type Calendar struct {
	events map[int]Event
	mu     sync.RWMutex
}

type Event struct {
	ID   int       `json:"id"`
	Name string    `json:"name"`
	Time time.Time `json:"time"`
}

type Result struct {
	Result []Event `json:"result"`
}

func SerializeEventSlice(events []Event) ([]byte, error) {
	data := Result{events}
	result, err := json.Marshal(data)
	return result, err
}

func NewCalendar() *Calendar {
	return &Calendar{make(map[int]Event), sync.RWMutex{}}
}

func NewEvent(Time time.Time, Name string) *Event { // Create event struct
	LastIdMutex.Lock()
	LastID++
	LastIdMutex.Unlock()
	return &Event{LastID, Name, Time}
}

func (c *Calendar) CreateEvent(event *Event) { // Add event to calendar
	c.mu.Lock()
	c.events[event.ID] = *event
	c.mu.Unlock()
}

func (c *Calendar) UpdateEvent(ID int, Time time.Time, Name string) error { // Add new event instead of old one
	c.mu.RLock()
	event, ok := c.events[ID]
	if !ok {
		c.mu.RUnlock()
		return errors.New("no such event")
	}
	c.mu.RUnlock()

	if !Time.IsZero() {
		event.Time = Time
	}
	if Name != "" {
		event.Name = Name
	}
	c.mu.Lock()
	c.events[ID] = event
	c.mu.Unlock()
	return nil
}

func (c *Calendar) DeleteEvent(ID int) (*Event, error) {
	c.mu.RLock()
	if _, ok := c.events[ID]; !ok {
		c.mu.RUnlock()
		return nil, errors.New("no such event")
	}
	c.mu.RUnlock()
	c.mu.Lock()
	deleted := c.events[ID]
	delete(c.events, ID)
	c.mu.Unlock()
	return &deleted, nil
}

func (c *Calendar) EventsForDay() []Event {
	var result []Event
	tYear, tMonth, tDay := time.Now().Date() // today
	c.mu.RLock()
	for _, v := range c.events {
		year, month, day := v.Time.Date()
		if tYear == year && tMonth == month && tDay == day {
			result = append(result, v)
		}
	}
	c.mu.RUnlock()
	return result
}

func (c *Calendar) EventsForWeek() []Event {
	var result []Event
	tYear, tWeek := time.Now().ISOWeek()
	c.mu.RLock()
	for _, v := range c.events {
		year, week := v.Time.ISOWeek()
		if tYear == year && tWeek == week {
			result = append(result, v)
		}
	}
	c.mu.RUnlock()
	return result
}

func (c *Calendar) EventsForMonth() []Event {
	var result []Event
	tYear, tMonth, _ := time.Now().Date() // today
	c.mu.RLock()
	for _, v := range c.events {
		year, month, _ := v.Time.Date()
		if tYear == year && tMonth == month {
			result = append(result, v)
		}
	}
	c.mu.RUnlock()
	return result
}

type CalendarHandler struct {
	calendar *Calendar
}

func NewCalendarHandler() *CalendarHandler {
	return &CalendarHandler{NewCalendar()}
}

func (c *CalendarHandler) CreateEventRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		SendResult(w, []byte("method not allowed"))
	}

	newEvent, err := ParseCreateRequest(r)
	if err != nil {
		SendError(w, err, http.StatusBadRequest)
		return
	}

	c.calendar.CreateEvent(newEvent)
	SendResult(w, []byte("new event created"))
}

func (c *CalendarHandler) UpdateEventRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		SendResult(w, []byte("method not allowed"))
	}

	ID, Time, Name, err := ParseUpdateRequest(r)
	if err != nil {
		SendError(w, err, http.StatusBadRequest)
		return
	}

	err = c.calendar.UpdateEvent(ID, Time, Name)
	if err != nil {
		SendError(w, err, http.StatusBadRequest)
		return
	}

	SendResult(w, []byte(fmt.Sprintf("event #%d updated", ID)))
}

func (c *CalendarHandler) DeleteEventRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		SendResult(w, []byte("method not allowed"))
	}

	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/x-www-form-urlencoded" {
		SendError(w, errors.New("invalid data"), http.StatusBadRequest)
		return
	}

	err := r.ParseForm()
	if err != nil {
		SendError(w, err, http.StatusBadRequest)
		return
	}

	ID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		SendError(w, err, http.StatusBadRequest)
		return
	}

	deleted, err := c.calendar.DeleteEvent(ID)
	if err != nil {
		SendError(w, err, http.StatusBadRequest)
		return
	}

	SendResult(w, []byte(fmt.Sprintf("event #%d (%s, %v) removed", deleted.ID, deleted.Name, deleted.Time)))
}

func (c *CalendarHandler) EventsForDayRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		SendResult(w, []byte("method not allowed"))
	}

	data, err := SerializeEventSlice(c.calendar.EventsForDay())
	if err != nil {
		SendError(w, err, http.StatusServiceUnavailable)
		return
	}

	SendResult(w, data)
}

func (c *CalendarHandler) EventsForWeekRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		SendResult(w, []byte("method not allowed"))
	}

	data, err := SerializeEventSlice(c.calendar.EventsForWeek())
	if err != nil {
		SendError(w, err, http.StatusServiceUnavailable)
		return
	}

	SendResult(w, data)
}

func (c *CalendarHandler) EventsForMonthRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		SendResult(w, []byte("method not allowed"))
	}

	data, err := SerializeEventSlice(c.calendar.EventsForMonth())
	if err != nil {
		SendError(w, err, http.StatusServiceUnavailable)
		return
	}

	SendResult(w, data)
}

type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type ResultResponse struct {
	Result []byte `json:"result"`
}

func ParseCreateRequest(r *http.Request) (*Event, error) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/x-www-form-urlencoded" {
		return nil, errors.New("invalid data")
	}
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}

	t, err := time.Parse("2006-01-02 15:04", r.FormValue("time"))
	if err != nil {
		return nil, err
	}

	name := r.FormValue("name")
	if name == "" {
		return nil, errors.New("name can't be blank")
	}

	newEvent := NewEvent(t, name)
	return newEvent, nil
}

func ParseUpdateRequest(r *http.Request) (int, time.Time, string, error) {
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/x-www-form-urlencoded" {
		return -1, time.Time{}, "", errors.New("invalid data")
	}
	err := r.ParseForm()
	if err != nil {
		return -1, time.Time{}, "", err
	}

	ID, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		return -1, time.Time{}, "", err
	}

	Time := r.FormValue("time")
	parsedTime, err := time.Parse("2006-01-02 15:04", r.FormValue("time"))
	if !(Time == "" || err == nil) {
		return -1, time.Time{}, "", err
	}

	Name := r.FormValue("name")

	return ID, parsedTime, Name, nil
}

func SendError(w http.ResponseWriter, err error, statusCode int) {
	data := ErrorResponse{err.Error()}
	result, _ := json.Marshal(data)
	w.WriteHeader(statusCode)
	w.Write(result)
}

func SendResult(w http.ResponseWriter, response []byte) {
	data := ResultResponse{response}
	result, _ := json.Marshal(data)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func main() {
	ch := NewCalendarHandler()
	mux := http.NewServeMux()
	mux.HandleFunc("/create_event", ch.CreateEventRequest)
	mux.HandleFunc("/update_event", ch.UpdateEventRequest)
	mux.HandleFunc("/delete_event", ch.DeleteEventRequest)
	mux.HandleFunc("/events_for_day", ch.EventsForDayRequest)
	mux.HandleFunc("/events_for_week", ch.EventsForWeekRequest)
	mux.HandleFunc("/events_for_month", ch.EventsForMonthRequest)

	wrappedMux := NewLogger(mux)

	http.ListenAndServe("localhost:8080", wrappedMux)
}
