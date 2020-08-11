package datastructure

// Enqueue 末尾にデータを追加する
func Enqueue(queue []int, number int) []int {
	queue = append(queue, number)
	return queue
}

// Dequeue 先頭にあるデータを取り出す
func Dequeue(queue []int) ([]int, int) {
	n := queue[0]
	queue = queue[1:]

	return queue, n
}
