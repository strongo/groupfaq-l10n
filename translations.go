package trans

import bots "github.com/strongo/bots-framework/core"

var TRANS = map[string]map[string]string{
	"EXAMPLE": {
		"en-US": "SAMPLE",
		"ru-RU": "ПРИМЕР",
		"it-IT": "ESEMPIO",
		"fa-IR": "نمونه",
	},
	bots.MESSAGE_TEXT_I_DID_NOT_UNDERSTAND_THE_COMMAND: {
		"en-US": "\xF0\x9F\x98\xB3 Sorry, I did not understand your order. May be I'm a little bit dumb...\n\nYou can return to main /menu",
		"ru-RU": "\xF0\x9F\x98\xB3 Извините, я не понял вашу команду. Возможно я немного туповат...\n\nЧтобы начать сначала нажмите /menu",
		"fa-IR": "\xF0\x9F\x98\xB3 ببخشید، من دستور شما را نفهمیدم. احتمالا کمی کند ذهن هستم...\n\nشما میتوانید به /منو ی اصلی بازگردید",
		"it-IT": "\xF0\x9F\x98\xB3 Scusami ma non ho capito cosa vuoi. Sono ancora un po' sciocco...\n\nPuoi ritornare al Menu con /menu",
	},
	RULES_AND_FAQ_of_GROPNAME: {
		"en-US": "<b>Rules & FAQ</b> of <b>%v</b>",
		"ru-RU": "<b>Правила и ЧаВо</b> для <b>%v</b>",
		"it-IT": "",
		"fa-IR": "",
	},
	PLEASE_ACCEPT_RULES: {
		"en-US": "Please accept rules.",
		"ru-RU": "Пожалуйста примите правила.",
		"it-IT": "",
		"fa-IR": "",
	},
	I_ACCEPT_RULES: {
		"en-US": "I accept rules of the group",
		"ru-RU": "Я принимаю правила группы",
		"it-IT": "",
		"fa-IR": "",
	},
	FAQ_TITLE: {
		"en-US": "<b>F</b>requently <b>A</b>sked <b>Q</b>uestions",
		"ru-RU": "<b>Ча</b>стые <b>Во</b>просы",
		"it-IT": "",
		"fa-IR": "",
	},
	PLEASE_READ_RULES_AND_FAQ: {
		"en-US": `Please read <a href="https://t.me/GroupFAQbot?start=joined_%v_%v">Rules & FAQ</a> before posting to this group.`,
		"ru-RU": `Пожалуйста прочитайте <a href="https://t.me/GroupFAQbot?start=joined_%v_%v">Правила и Частые вопросы</a> прежде чем писать в группу.`,
		"it-IT": "",
		"fa-IR": "",
	},
	USERNAME_JOINED_GROUP: {
		"en-US": "<b>%v</b> joined group.",
		"ru-RU": "<b>%v</b> присоеденился к группе.",
		"it-IT": "",
		"fa-IR": "",
	},
	NUMBER_OF_JOINED_MEMBERS: {
		"en-US": "%d new members:",
		"ru-RU": "%d новых участников:",
		"it-IT": "",
		"fa-IR": "",
	},
	"": {
		"en-US": "",
		"ru-RU": "",
		"it-IT": "",
		"fa-IR": "",
	},
}
