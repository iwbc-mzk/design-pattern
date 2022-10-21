package main

func main() {
	table1 := NewTable([]string{"名前", "年齢", "出身地"}...)
	row1 := NewRow([]string{"田中", "22歳", "東京"}...)

	table2 := NewTable([]string{"佐藤", "54歳", "北海道"}...)
	row3 := NewRow([]string{"168cm", "59kg", "A型"}...)

	table3 := NewTable([]string{"北村", "20歳", "千葉"}...)
	row4 := NewRow([]string{"189cm", "90kg", "O型"}...)
	row5 := NewRow([]string{"会社員", "未婚"}...)
	table4 := NewTable("銀行員として勤務中。趣味は登山。")
	row6 := NewRow("明日は登山会の仲間と富士山に登山予定。")

	table2.add(row3)
	table4.add(row6)
	table3.add(row4, row5, table4)
	table1.add(row1, table2, table3)
	
	table1.show(0)
}