package utils

import (
	"CLI_TodoApp/internal/todo"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

func ParseId(s string) (int, error) {
	return strconv.Atoi(s)
}

// In danh sách công việc dạng bảng
func PrintList(items []todo.Item) {
	if len(items) == 0 {
		fmt.Println("📭 Danh sách đang trống. Hãy thêm công việc đầu tiên!")
		return
	}

	fmt.Println("")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	fmt.Fprintln(w, "ID\tTRẠNG THÁI\tNỘI DUNG CÔNG VIỆC\tNGÀY TẠO")

	for _, item := range items {
		status := "[ ] Đang làm"
		if item.Done {
			status = "[✔] Hoàn thành"
		}
		dateStr := item.CreatedAt.Format("02/01 15:04")
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", item.Id, status, item.Task, dateStr)
	}
	w.Flush()
	fmt.Println("")
}

// In hướng dẫn sử dụng chi tiết (Cheat Sheet)
func PrintHelp() {
	fmt.Println("\n================ HƯỚNG DẪN SỬ DỤNG ================")
	fmt.Println("🔹 Thêm công việc mới:")
	fmt.Println("   Gõ: add <nội dung>  hoặc  a <nội dung>")
	fmt.Println("   Ví dụ: a Đi mua sữa")

	fmt.Println("\n🔹 Xem danh sách:")
	fmt.Println("   Gõ: list  hoặc  ls")

	fmt.Println("\n🔹 Đánh dấu đã làm xong:")
	fmt.Println("   Gõ: done <id>  hoặc  d <id>")
	fmt.Println("   Ví dụ: d 1 (Để xong việc số 1)")

	fmt.Println("\n🔹 Xóa công việc:")
	fmt.Println("   Gõ: del <id>  hoặc  rm <id>")
	fmt.Println("   Ví dụ: rm 2 (Để xóa việc số 2)")

	fmt.Println("\n🔹 Thoát chương trình:")
	fmt.Println("   Gõ: exit  hoặc  q")
	fmt.Println("====================================================")
}
