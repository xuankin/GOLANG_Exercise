package main

import (
	"CLI_TodoApp/cmd/utils"
	"CLI_TodoApp/internal/storage"
	"CLI_TodoApp/internal/todo"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1. Cài đặt nơi lưu file
	todoFilePath := "todos.json"
	s := storage.NewJSONStore(todoFilePath)
	list, err := todo.NewTodoList(s)
	if err != nil {
		fmt.Printf("Lỗi khởi động: %v\n", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	// 2. Màn hình chào mừng thân thiện
	fmt.Println("\n👋 CHÀO MỪNG BẠN ĐẾN VỚI TODO CLI!")
	fmt.Println("Tip: Gõ 'help' hoặc 'h' để xem hướng dẫn chi tiết.")

	// Hiện danh sách ngay khi vào để người dùng nắm tình hình
	utils.PrintList(list.Items)

	// 3. Vòng lặp chính
	for {
		// Dấu nhắc lệnh
		fmt.Print("\n(Nhập lệnh) > ")

		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		command := strings.ToLower(parts[0]) // Chuyển về chữ thường để a hay A đều được
		args := parts[1:]

		switch command {
		case "exit", "quit", "q":
			fmt.Println("👋 Tạm biệt! Hẹn gặp lại.")
			return

		case "help", "h":
			utils.PrintHelp()

		case "list", "ls", "l":
			utils.PrintList(list.Items)

		case "add", "a":
			var text string
			if len(args) == 0 {
				// HƯỚNG DẪN KHI NGƯỜI DÙNG QUÊN NHẬP NỘI DUNG
				fmt.Print("✍️  Bạn muốn làm gì? (Nhập nội dung rồi Enter): ")
				if scanner.Scan() {
					text = scanner.Text()
				}
			} else {
				text = strings.Join(args, " ")
			}

			if strings.TrimSpace(text) == "" {
				fmt.Println("⚠️  Bạn chưa nhập nội dung, hủy thao tác.")
				continue
			}

			item, _ := list.Add(text)
			fmt.Printf("✅ Đã thêm việc mới: [%d] %s\n", item.Id, item.Task)

		case "done", "d":
			if len(args) == 0 {
				// HƯỚNG DẪN KHI QUÊN ID
				fmt.Print("🎉 Bạn đã xong việc số mấy? (Nhập ID): ")
				if scanner.Scan() {
					args = append(args, scanner.Text())
				}
			}
			if len(args) == 0 {
				continue
			} // Nếu vẫn không nhập gì thì thôi

			id, err := utils.ParseId(args[0])
			if err != nil {
				fmt.Println("⚠️  ID phải là một con số (Ví dụ: 1, 2).")
				continue
			}
			if err := list.MarkDone(id); err != nil {
				fmt.Printf("❌ Lỗi: %v\n", err)
			} else {
				fmt.Printf("👏 Tuyệt vời! Đã hoàn thành công việc [%d]\n", id)
			}

		case "delete", "del", "rm":
			if len(args) == 0 {
				// HƯỚNG DẪN KHI QUÊN ID
				fmt.Print("🗑️  Bạn muốn xóa việc số mấy? (Nhập ID): ")
				if scanner.Scan() {
					args = append(args, scanner.Text())
				}
			}
			if len(args) == 0 {
				continue
			}

			id, err := utils.ParseId(args[0])
			if err != nil {
				fmt.Println("⚠️  ID phải là số.")
				continue
			}
			if err := list.Delete(id); err != nil {
				fmt.Printf("❌ Lỗi: %v\n", err)
			} else {
				fmt.Printf("🗑️  Đã xóa vĩnh viễn công việc [%d]\n", id)
			}

		case "clear":
			fmt.Print("⚠️  Bạn có chắc muốn xóa HẾT danh sách không? (y/n): ")
			if scanner.Scan() && strings.ToLower(scanner.Text()) == "y" {
				list.Clear()
				fmt.Println("🧹 Đã dọn sạch danh sách.")
			} else {
				fmt.Println("Đã hủy lệnh xóa.")
			}

		default:
			fmt.Printf("🤔 Không hiểu lệnh '%s'. Gõ 'help' để xem hướng dẫn nhé.\n", command)
		}
	}
}
