package test

import (
	"log"
	_ "novel_crawler/bootstrap"
	"novel_crawler/global/variable"
	"regexp"
	"testing"
)

func Test1(t *testing.T) {

	str := `
		<div    id="container">
			12
			<p>13</p>
			<span>24</span>
			<a href="123"></a>
			<br>
			<br/>
			<br />
		</div>
	`

	baseReg := variable.InfoStore.GetBaseRegReplace()

	for k, v := range baseReg {
		reg := regexp.MustCompile(k)
		str = reg.ReplaceAllString(str, v)
	}
	log.Println(str)
}
