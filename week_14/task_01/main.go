package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

const dataFile = "notes.json"

// Note представляет заметку
type Note struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

// NotesManager управляет заметками
type NotesManager struct {
	Notes  []Note `json:"notes"`
	NextID int    `json:"next_id"`
}

// Load загружает заметки из файла
func (nm *NotesManager) Load() error {
	data, err := os.ReadFile(dataFile)
	if os.IsNotExist(err) {
		nm.NextID = 1
		return nil
	}
	if err != nil {
		return err
	}
	return json.Unmarshal(data, nm)
}

// Save сохраняет заметки в файл
func (nm *NotesManager) Save() error {
	data, err := json.MarshalIndent(nm, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}

// Add добавляет новую заметку
func (nm *NotesManager) Add(text string) {
	note := Note{
		ID:        nm.NextID,
		Text:      text,
		CreatedAt: time.Now(),
	}
	nm.Notes = append(nm.Notes, note)
	nm.NextID++
	fmt.Printf("Заметка #%d добавлена\n", note.ID)
}

// List выводит все заметки
func (nm *NotesManager) List() {
	if len(nm.Notes) == 0 {
		fmt.Println("Заметок нет")
		return
	}
	fmt.Println("=== Заметки ===")
	for _, n := range nm.Notes {
		fmt.Printf("#%d [%s]: %s\n", n.ID, n.CreatedAt.Format("02.01.2006 15:04"), n.Text)
	}
}

// Delete удаляет заметку по ID
func (nm *NotesManager) Delete(id int) {
	for i, n := range nm.Notes {
		if n.ID == id {
			nm.Notes = append(nm.Notes[:i], nm.Notes[i+1:]...)
			fmt.Printf("Заметка #%d удалена\n", id)
			return
		}
	}
	fmt.Printf("Заметка #%d не найдена\n", id)
}

// Search ищет заметки по тексту
func (nm *NotesManager) Search(query string) {
	query = strings.ToLower(query)
	found := false
	for _, n := range nm.Notes {
		if strings.Contains(strings.ToLower(n.Text), query) {
			fmt.Printf("#%d: %s\n", n.ID, n.Text)
			found = true
		}
	}
	if !found {
		fmt.Println("Ничего не найдено")
	}
}

func main() {
	addFlag := flag.String("add", "", "Добавить заметку")
	listFlag := flag.Bool("list", false, "Показать все заметки")
	deleteFlag := flag.Int("delete", 0, "Удалить заметку по ID")
	searchFlag := flag.String("search", "", "Поиск заметок")
	flag.Parse()

	nm := &NotesManager{}
	if err := nm.Load(); err != nil {
		fmt.Println("Ошибка загрузки:", err)
		return
	}

	switch {
	case *addFlag != "":
		nm.Add(*addFlag)
	case *listFlag:
		nm.List()
	case *deleteFlag > 0:
		nm.Delete(*deleteFlag)
	case *searchFlag != "":
		nm.Search(*searchFlag)
	default:
		fmt.Println("Используйте: -add, -list, -delete, -search")
		flag.PrintDefaults()
	}

	if err := nm.Save(); err != nil {
		fmt.Println("Ошибка сохранения:", err)
	}
}

