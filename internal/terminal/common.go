package terminal

import log "github.com/sirupsen/logrus"

func judeWidth(terminalWidth int, fieldSum int, fieldImportance []string, fieldMinWidth []int, delField []string, i int) []string {
	if terminalWidth >= fieldSum {
		return delField
	} else {
		fieldSum -= fieldMinWidth[i]
		terminalWidth += 1
		delField = append(delField, fieldImportance[i])
		i++
		return judeWidth(terminalWidth, fieldSum, fieldImportance, fieldMinWidth, delField, i)
	}
}

func deleteArray(array []string, delNames []string) []string {
	ret := make([]string, 0, len(array))
	for _, v := range array {
		flag := true
		for _, delName := range delNames {
			if v == delName {
				flag = false
			}
		}
		if flag == true {
			ret = append(ret, v)
		}
	}
	return ret
}

func tableFieldAdap(fieldOrder []string, fieldImportance []string, fieldMinWidth []int, terminalWidth int) []string {
	// fieldImportance和fieldWidth顺序对应, 重要性逐渐下降
	fieldNum := len(fieldOrder)
	log.Info("source: ", fieldOrder, terminalWidth)
	terminalWidth = terminalWidth - (fieldNum - 1) // 字段以空格为间隔

	fieldSum := 0
	for _, v := range fieldMinWidth {
		fieldSum += v
	}

	i := 0 // 次数
	delField := judeWidth(terminalWidth, fieldSum, fieldImportance, fieldMinWidth, []string{}, i)

	fieldOrder = deleteArray(fieldOrder, delField)
	log.Info("fieldOrder: ", fieldOrder)
	return fieldOrder
}
