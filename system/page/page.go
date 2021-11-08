package page

import (
	"html/template"
	"math"
	"strconv"
)

type page struct {
	Pre      string //上一页样式
	Next     string //下一页样式
	Count    int    //总数
	NowPage  int    //当前页
	PageSize int    //一页条数
	Current  int    //当前页
	PrePage  int    //上一页
	NextPage int    //下一页
	MaxNum   int    //页数按钮显示多少个
}

func GetPagination(nowPage, pageSize, count int, param string) template.HTML {
	var p = &page{
		Pre:    "上一页",
		Next:   "下一页",
		MaxNum: 10,
	}

	return p.Pagination(nowPage,pageSize,count,param)
}

func (p *page) Pagination(nowPage, pageSize, count int, param string) template.HTML {
	if param == "" {
		param = "uuid=sblog"
	}
	p.NowPage = nowPage
	p.Count = count
	if nowPage <= 1 {
		nowPage = 1
	}
	//总页数
	num := float64(count) / float64(pageSize)
	pageCount := int(math.Ceil(num))

	if count <= 1 || pageCount <= 1 {
		//<ul class="pagination"><li class="disabled"><span>` + p.Pre + `</span></li><li class="active"><span>1</span></li><li class="disabled"><span>` + p.Next + `</span></li></ul>
		return template.HTML(`<div class="pagelist"><a href="javascript:;">` + p.Pre + `</a>&nbsp;<b>1</b>&nbsp;<a href="javascript:;">` + p.Next + `</a>&nbsp;</div>`)
	}

	p.Current = nowPage //当前页
	p.Count = pageCount //总页数

	if nowPage <= 1 {
		p.PrePage = 1 //上一页
	} else {
		p.PrePage = nowPage - 1 //上一页
	}

	if nowPage >= pageCount {
		p.NextPage = pageCount //下一页
	} else {
		p.NextPage = nowPage + 1 //下一页
	}

	var htmls string
	//拼接上面的
	htmls = `<div class="pagelist">`
	if p.Current <= 1 {
		htmls += `<a href="javascript:;">` + p.Pre + `</a>&nbsp;`
	} else {
		htmls += `<a href="?p=` + strconv.Itoa(p.PrePage) + `">` + p.Pre + `</a>&nbsp;`
	}

	htmls += middleStr(p.Current, pageCount, p.MaxNum, param)

	//拼接底页
	if p.Current >= pageCount {
		htmls += `<a href="javascript:;">` + p.Next + `</a>&nbsp;`
	} else {
		htmls += `<a href="?p=` + strconv.Itoa(p.NextPage) + `">` + p.Next + `</a>&nbsp;`
	}

	htmls += `</div>`

	return template.HTML(htmls)
}

//中间部分拼接
func middleStr(current, pageCount, maxCount int, param string) string {
	var htmls string
	var i int
	//如果总页数不是很多就全部展示出来
	if pageCount <= maxCount {
		for i = 1; i <= pageCount; i++ {
			if current == i {
				htmls += `<b>` + strconv.Itoa(i) + `</b>&nbsp;`
			} else {
				if param != "" {
					htmls += `<a href="?p=` + strconv.Itoa(i) + `">` + strconv.Itoa(i) + `</a>&nbsp;`
				} else {
					htmls += `<a href="?` + param + `&p=` + strconv.Itoa(i) + `">` + strconv.Itoa(i) + `</a>&nbsp;`
				}

			}
		}
		return htmls
	}
	//如果分页比较多就展示部分按钮，其他部分用...代替
	if current > 1 {
		htmls += `<a href="?p=1">...</a>&nbsp;`
	}
	//获取最大值
	maxCount -= 1
	now := maxCount + current
	if now >= pageCount {
		now = pageCount
	}
	//里面必须有特定个数的按钮
	cnum := now - maxCount
	for i = cnum; i <= now; i++ {
		if current == i {
			htmls += `<b>` + strconv.Itoa(i) + `</b>&nbsp;`
		} else {
			if param != "" {
				htmls += `<a href="?p=` + strconv.Itoa(i) + `">` + strconv.Itoa(i) + `</a>&nbsp;`
			} else {
				htmls += `<a href="?` + param + `&p=` + strconv.Itoa(i) + `">` + strconv.Itoa(i) + `</a>&nbsp;`
			}
		}
	}
	//...结尾
	if now < pageCount && current >= 1 {
		htmls += `<a href="?p=` + strconv.Itoa(now) + `">...</a>&nbsp;`
	}
	return htmls
}
