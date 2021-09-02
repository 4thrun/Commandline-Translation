package config

// Using data from 3 sites
// 1. iCIBA: http://open.iciba.com/index.php
// 2. Youdao: http://fanyi.youdao.com/openapi
// 3. Google: http://translate.google.cn/translate_a/single
type myAPI struct {
	Iciba  string `json:"iciba"`
	Youdao string `json:"youdao"`
	Google string `json:"google"`
}

var API myAPI

func init() {
	API = myAPI{
		// iCIBA: http://open.iciba.com/ is now 403 ( ¯(∞)¯ )
		// for iCIBA, do use XML format (JSON has less information)
		Iciba: "http://dict-co.iciba.com/api/dictionary.php?key=D191EBD014295E913574E1EAF8E06666&type=xml&w=${word}",
		// Youdao: 1000 requests per hour at most
		Youdao: "http://fanyi.youdao.com/openapi.do?keyfrom=node-fanyi&key=110811608&type=data&doctype=json&version=1.1&q=${word}",
		// Google: generous
		Google: "http://translate.google.cn/translate_a/single?client=gtx&dt=t&dj=1&ie=UTF-8&sl=auto&tl=zh_CN&q=${word}",
	}
}
