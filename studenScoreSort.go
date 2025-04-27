package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

//** 有如下结构的json 字符串, 根据成绩输出每个班级学生的名次.

//  "{\"一班\":{\"张小丙\":87,\"张小甲\":98,\"张小乙\":90},\"二班\":{\"王七六\":76,\"王九七\":97,\"胡八一\":81,\"王六零\":60,\"刘八一\":81,\"李八一\":81}}"
// {
//   "一班": {
//     "张小丙": 87,
//     "张小甲": 98,
//     "张小乙": 90
//   },
//   "二班": {
//     "王七六": 76,
//     "王九七": 97,
//     "胡八一": 81,
//     "王六零": 60,
//     "刘八一": 81,
//     "李八一": 81
//   }
// }
// 题目解释
// 每个班级分别排名.

// 分数相同的时候名次相同.

// 当出现相同分数的情况下, 名次并不连续. 既排名在两个并列第一之后的学生名次是第三, 排名在三个并列第一之后的学生名次是第四.

// 输出示例(不需要考虑输出顺序):

// 一班 张小丙 第3名
// 一班 张小甲 第1名
// 一班 张小乙 第2名

// 二班 王七六 第5名
// 二班 王九七 第1名
// 二班 胡八一 第2名
// 二班 王六零 第6名
// 二班 刘八一 第2名
// 二班 李八一 第2名**//
// 成绩排序
func main9() {
	// JSON字符串
	class2 := `{
  "一班": {
    "张小丙": 87,
    "张小甲": 98,
    "张小乙": 90
  },
  "二班": {
    "王七六": 76,
    "王九七": 97,
    "胡八一": 81,
    "王六零": 60,
    "刘八一": 81,
    "李八一": 81
  }
}`

	// 将JSON字符串解析为map
	classmap := make(map[string]map[string]int)
	err := json.Unmarshal([]byte(class2), &classmap)
	if err != nil {
		fmt.Println("解析JSON出错:", err)
		return
	}
	for classname, students := range classmap {
		type student struct {
			name  string
			score int
		}
		var studentscores []student

		for name, score := range students {
			studentscores = append(studentscores, student{name: name, score: score})
		}

		sort.Slice(studentscores, func(i, j int) bool {
			return studentscores[i].score > studentscores[j].score
		})
		rankMap := make(map[string]int)
		currentRank := 1
		prevScore := -1
		skipCount := 0

		for i, student := range studentscores {
			if i > 0 && student.score == prevScore {
				rankMap[student.name] = rankMap[studentscores[i-1].name]
				skipCount++
			} else {
				rankMap[student.name] = currentRank + skipCount
				currentRank++
				skipCount = 0
			}
			prevScore = student.score
		}
		for _, student := range studentscores {
			fmt.Printf("%s %s 第%d名\n", classname, student.name, rankMap[student.name])
		}
	}
}
