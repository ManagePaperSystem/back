package models

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"gopkg.in/gomail.v2"
	"math"
	"math/big"
	"strconv"
)

var (
	Limit     int
	TimeStr   string
	DB        *sql.DB
	PhaseFlag int
	Root      TreeNode
)

var Operator = []string{
	"+", "-", "*", "/",
	"sqrt{", "^2",
	"sin", "cos", "tan",
}

var triangle = []float64{
	0, 30, 45, 60, 90,
	120, 135, 150, 180,
	210, 225, 240, 270,
	300, 315, 330, 360,
}

type Question struct {
	PaperName string
	Question  string
	Answer    string
	A         string
	B         string
	C         string
	D         string
	Choice    string
}

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   string
	Level int
}

func SendEmail(e string) (passcode string, flag bool) {
	var mailConf MailboxConf
	mailConf.Title = "中小学作业系统验证码"
	mailConf.RecipientList = []string{e}
	mailConf.Sender = `281597094@qq.com`
	mailConf.SPassword = "ysdskdgrqgazbiag"
	mailConf.SMTPAddr = `smtp.qq.com`
	mailConf.SMTPPort = 25
	rnd := RandNumber(1000000)
	passcode = fmt.Sprintf("%06v", rnd)
	html := passcode
	flag = true
	m := gomail.NewMessage()
	m.SetHeader(`From`, mailConf.Sender, "作业系统官方")
	m.SetHeader(`To`, mailConf.RecipientList...)
	m.SetHeader(`Subject`, mailConf.Title)
	m.SetBody(`text/html`, html)
	err := gomail.NewDialer(mailConf.SMTPAddr, mailConf.SMTPPort, mailConf.Sender, mailConf.SPassword).DialAndSend(m)
	if err != nil {
		//log.Fatalf("Send Email Fail, %s", err.Error())
		flag = false
	}
	return
}

func InitTree(node *TreeNode, cnt int) *TreeNode {
	// if cnt == 0, the node stands for number
	if cnt == 0 {
		num := RandNumber(100) + 1
		node.Val, node.Left, node.Right, node.Level = strconv.Itoa(num), nil, nil, -1
		return node
	}
	// else, this node stands for operator
	num := RandNumber(Limit)
	node.Level = num
	node.Val = Operator[num]

	switch node.Level {
	case 4:
		num = 0
		PhaseFlag = Max(PhaseFlag, 6)
	case 5:
		num = cnt - 1
		PhaseFlag = Max(PhaseFlag, 6)
	case 6, 7, 8:
		PhaseFlag = Max(PhaseFlag, 9)
	default:
		num = RandNumber(cnt)
	}
	node.Left, node.Right = new(TreeNode), new(TreeNode)
	if node.Level < 6 {
		node.Left, node.Right = InitTree(node.Left, num), InitTree(node.Right, cnt-num-1)
	} else {
		// triangle have their own function calls InitTri
		node.Left, node.Right = InitTree(node.Left, 0), InitTri(node.Right, cnt-1)
	}
	return node
}

// InitTri will build a son-tree promising triangle function will operator a special number
// The special number will be not a long expression but a number
// which gives a more suitable expression
func InitTri(node *TreeNode, cnt int) *TreeNode {
	if cnt == 0 {
		num := RandNumber(17)
		node.Val, node.Left, node.Right, node.Level = strconv.FormatFloat(triangle[num], 'f', 2, 64), nil, nil, -1
		return node
	}
	num := RandNumber(4)
	node.Level = num
	node.Val = Operator[num]
	node.Left, node.Right = new(TreeNode), new(TreeNode)
	node.Left, node.Right = InitTri(node.Left, 0), InitTree(node.Right, cnt-1)
	return node
}

// TransferTree organize an expression with In-order method
// each Answer is {LeftAnswer Operator RightAnswer}, be careful of the operator order
// In the end, Formula, Answer, flag are the complete expression, the answer of the expression, the accuracy of the expression
func TransferTree(node *TreeNode) (Formula, Answer string, flag bool) {
	if node.Level == -1 {
		Formula = node.Val
		Answer = Formula
		flag = true
		return
	}
	LeftFormula, LeftAnswer, LeftFlag := TransferTree(node.Left)
	RightFormula, RightAnswer, RightFlag := TransferTree(node.Right)
	flag = LeftFlag && RightFlag
	switch node.Level {
	case 0, 1:
		if node.Left.Level <= 1 && node.Left.Level >= 0 {
			LeftFormula = "(" + LeftFormula + ")"
		}
		Formula = LeftFormula + " " + node.Val + " " + RightFormula
		Answer = LeftAnswer + " " + node.Val + " " + RightAnswer
		num1, _ := strconv.ParseFloat(LeftAnswer, 64)
		num2, _ := strconv.ParseFloat(RightAnswer, 64)
		if node.Level == 0 {
			Answer = strconv.FormatFloat(num1+num2, 'f', 2, 64)
		} else {
			Answer = strconv.FormatFloat(num1-num2, 'f', 2, 64)
		}
	case 2, 3:
		if node.Left.Level <= 1 && node.Left.Level >= 0 {
			LeftFormula = "(" + LeftFormula + ")"
		}
		if node.Right.Level >= 0 && node.Right.Level < 4 {
			RightFormula = "(" + RightFormula + ")"
		}
		Formula = LeftFormula + " " + node.Val + " " + RightFormula
		num1, _ := strconv.ParseFloat(LeftAnswer, 64)
		num2, _ := strconv.ParseFloat(RightAnswer, 64)
		if node.Level == 2 {
			Answer = strconv.FormatFloat(num1*num2, 'f', 2, 64)
		} else {
			Answer = strconv.FormatFloat(num1/num2, 'f', 2, 64)
		}
	case 4:
		Formula = `\` + node.Val + RightFormula + "}"
		num, _ := strconv.ParseFloat(RightAnswer, 64)
		if num < 0 {
			flag = false
			return
		}
		Answer = strconv.FormatFloat(math.Sqrt(num), 'f', 2, 64)
	case 5:
		if node.Left.Level != -1 {
			LeftFormula = "(" + LeftFormula + ")"
		}
		Formula = LeftFormula + node.Val
		num, _ := strconv.ParseFloat(LeftAnswer, 64)
		Answer = strconv.FormatFloat(num*num, 'f', 2, 64)
	default:
		Formula = node.Val + RightFormula
		if node.Right.Level == -1 {
			num, _ := strconv.ParseFloat(RightAnswer, 64)
			num = num / 180 * math.Pi
			if node.Level == 6 {
				Answer = strconv.FormatFloat(math.Sin(num), 'f', 2, 64)
			} else if node.Level == 7 {
				Answer = strconv.FormatFloat(math.Cos(num), 'f', 2, 64)
			} else {
				if int(num)%90 == 0 {
					flag = false
					return
				}
				Answer = strconv.FormatFloat(math.Tan(num), 'f', 2, 64)
			}
		} else {
			num, _ := strconv.ParseFloat(node.Right.Left.Val, 64)
			num1 := num / 180 * math.Pi
			ra, _ := strconv.ParseFloat(RightAnswer, 64)

			if node.Right.Level == 0 {
				n := ra - num
				// num + n = ra
				if node.Level == 6 {
					Answer = strconv.FormatFloat(math.Sin(num1)+n, 'f', 2, 64)
				} else if node.Level == 7 {
					Answer = strconv.FormatFloat(math.Cos(num1)+n, 'f', 2, 64)
				} else {
					if int(num)%90 == 0 {
						flag = false
						return
					}
					Answer = strconv.FormatFloat(math.Tan(num1)+n, 'f', 2, 64)
				}
			} else if node.Right.Level == 1 {
				n := num - ra
				// num - n = ra
				if node.Level == 6 {
					Answer = strconv.FormatFloat(math.Sin(num1)-n, 'f', 2, 64)
				} else if node.Level == 7 {
					Answer = strconv.FormatFloat(math.Cos(num1)-n, 'f', 2, 64)
				} else {
					if int(num)%90 == 0 {
						flag = false
						return
					}
					Answer = strconv.FormatFloat(math.Tan(num1)-n, 'f', 2, 64)
				}
			} else if node.Right.Level == 2 {
				n := ra / num
				// num * n = ra
				if node.Level == 6 {
					Answer = strconv.FormatFloat(math.Sin(num1)*n, 'f', 2, 64)
				} else if node.Level == 7 {
					Answer = strconv.FormatFloat(math.Cos(num1)*n, 'f', 2, 64)
				} else {
					if int(num)%90 == 0 {
						flag = false
						return
					}
					Answer = strconv.FormatFloat(math.Tan(num1)*n, 'f', 2, 64)
				}
			} else {
				n := num / ra
				// num / n = ra
				if node.Level == 6 {
					Answer = strconv.FormatFloat(math.Sin(num1)/n, 'f', 2, 64)
				} else if node.Level == 7 {
					Answer = strconv.FormatFloat(math.Cos(num1)/n, 'f', 2, 64)
				} else {
					if int(num)%90 == 0 {
						flag = false
						return
					}
					Answer = strconv.FormatFloat(math.Tan(num1)/n, 'f', 2, 64)
				}
			}
		}
	}
	return
}

func RandChoice(ans string, c *[4]string) {
	c[0], c[1], c[2], c[3] = "", "", "", ""
	c[RandNumber(4)] = ans
	length := len(ans)
	for i := 0; i < 4; i++ {
		if c[i] != "" {
			continue
		}
		str := []byte(ans)
		Times := RandNumber(length-1) + 1
		for j := 0; j < Times; j++ {
			position := RandNumber(length)
			if str[position] == '.' || str[position] == '-' {
				j--
				continue
			}
			str[position] = uint8(RandNumber(10) + 48)
		}
		if string(str) == ans {
			i--
			continue
		}
		c[i] = string(str)
	}
}

func RandNumber(L int) int {
	x, _ := rand.Int(rand.Reader, big.NewInt(int64(L)))
	return int(x.Int64())
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
