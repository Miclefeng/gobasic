package parser

import (
	"code/CrawlerSingle/engine"
	"code/CrawlerSingle/model"
	"regexp"
	"strconv"
)

var (
	nameCompile      = regexp.MustCompile(`<th>[^<]*<a href="http[s]?://album.zhenai.com/u/[\d]+"[^>]*>([^<]+)</a>[^<]*</th>`)
	genderCompile    = regexp.MustCompile(`<td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)
	huKouCompile     = regexp.MustCompile(`<td><span class="grayL">居住地：</span>([^<]+)</td>`)
	ageCompile       = regexp.MustCompile(`<td width="180"><span class="grayL">年龄：</span>([^<]+)</td>`)
	educationCompile = regexp.MustCompile(`<td><span class="grayL">学.*?历：</span>([^<]+)</td>`)
	marriageCompile  = regexp.MustCompile(`<td width="180"><span class="grayL">婚况：</span>([^<]+)</td>`)
	heightCompile    = regexp.MustCompile(`<td width="180"><span class="grayL">身.*?高：</span>([^<]+)</td>`)
	incomeCompile    = regexp.MustCompile(`<td><span class="grayL">月.*?薪：</span>([^<]+)</td>`)
)

func ParseProfile(contents []byte) engine.ParserResult {

	var profiles []model.Profile
	var profile model.Profile
	var result engine.ParserResult

	names := extract(contents, nameCompile)
	for _, name := range names {
		profile = model.Profile{
			Name: name,
		}
		profiles = append(profiles, profile)
	}

	genders := extract(contents, genderCompile)
	for i, gender := range genders {
		profiles[i].Gender = gender
	}

	ages := extract(contents, ageCompile)
	for i, age := range ages {
		age, err := strconv.Atoi(age)
		if err != nil {
			age = 0
		}
		profiles[i].Age = age
	}

	heights := extract(contents, heightCompile)
	for i, height := range heights {
		height, err := strconv.Atoi(height)
		if err != nil {
			height = 0
		}
		profiles[i].Height = height
	}

	huKous := extract(contents, huKouCompile)
	for i, huKou := range huKous {
		profiles[i].HuKou = huKou
	}

	educations := extract(contents, educationCompile)
	for i, education := range educations {
		profiles[i].Education = education
	}

	marriages := extract(contents, marriageCompile)
	for i, marriage := range marriages {
		profiles[i].Marriage = marriage
	}

	incomes := extract(contents, incomeCompile)
	for i, income := range incomes {
		profiles[i].Income = income
	}

	for _, pf := range profiles {
		result.Items = append(result.Items, pf)
	}
	return result
}

func extract(contents []byte, re *regexp.Regexp) []string {
	var result []string
	matches := re.FindAllSubmatch(contents, -1)
	for _, match := range matches {
		if len(match) >= 2 {
			result = append(result, string(match[1]))
		}
	}

	return result
}
