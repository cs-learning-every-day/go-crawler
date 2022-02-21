package model

type Book struct {
	Title        string
	Author       string
	CoverImgUrl  string
	Grade        string
	WordCount    string
	ChapterCount string
	State        string
	UpdateTime   string
	Source       string // 来源
	Intro        string
}

//func (b Book) String() string {
//	return fmt.Sprintf("Book: {Title: %s}", b.Title)
//}
