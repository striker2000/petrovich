// DO NOT EDIT!
// Code generated from rules/rules.json

package petrovich

var allRules = rules{
	firstName: rulesGroup{
		exceptions: []rule{
			{
				gender: Male,
				test: []string{
					"лев",
				},
				mods: []string{
					"--ьва",
					"--ьву",
					"--ьва",
					"--ьвом",
					"--ьве",
				},
			},
			{
				gender: Male,
				test: []string{
					"пётр",
				},
				mods: []string{
					"---етра",
					"---етру",
					"---етра",
					"---етром",
					"---етре",
				},
			},
			{
				gender: Male,
				test: []string{
					"павел",
				},
				mods: []string{
					"--ла",
					"--лу",
					"--ла",
					"--лом",
					"--ле",
				},
			},
			{
				gender: Male,
				test: []string{
					"яша",
				},
				mods: []string{
					"-и",
					"-е",
					"-у",
					"-ей",
					"-е",
				},
			},
			{
				gender: Male,
				test: []string{
					"илья",
				},
				mods: []string{
					"-и",
					"-е",
					"-ю",
					"-ёй",
					"-е",
				},
			},
			{
				gender: Male,
				test: []string{
					"шота",
				},
				mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
			{
				gender: Female,
				test: []string{
					"агидель",
					"жизель",
					"нинель",
					"рашель",
					"рахиль",
				},
				mods: []string{
					"-и",
					"-и",
					".",
					"ю",
					"-и",
				},
			},
		},
		suffixes: []rule{
			{
				gender: Androgynous,
				test: []string{
					"е",
					"ё",
					"и",
					"о",
					"у",
					"ы",
					"э",
					"ю",
				},
				mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
			{
				gender: Male,
				test: []string{
					"уа",
					"иа",
				},
				mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
			{
				gender: Female,
				test: []string{
					"б",
					"в",
					"г",
					"д",
					"ж",
					"з",
					"й",
					"к",
					"л",
					"м",
					"н",
					"п",
					"р",
					"с",
					"т",
					"ф",
					"х",
					"ц",
					"ч",
					"ш",
					"щ",
					"ъ",
					"иа",
					"ль",
				},
				mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
			{
				gender: Female,
				test: []string{
					"ь",
				},
				mods: []string{
					"-и",
					"-и",
					".",
					"ю",
					"-и",
				},
			},
			{
				gender: Male,
				test: []string{
					"ь",
				},
				mods: []string{
					"-я",
					"-ю",
					"-я",
					"-ем",
					"-е",
				},
			},
			{
				gender: Androgynous,
				test: []string{
					"га",
					"ка",
					"ха",
					"ча",
					"ща",
					"жа",
				},
				mods: []string{
					"-и",
					"-е",
					"-у",
					"-ой",
					"-е",
				},
			},
			{
				gender: Female,
				test: []string{
					"ша",
				},
				mods: []string{
					"-и",
					"-е",
					"-у",
					"-ей",
					"-е",
				},
			},
			{
				gender: Male,
				test: []string{
					"ша",
					"ча",
					"жа",
				},
				mods: []string{
					"-и",
					"-е",
					"-у",
					"-ей",
					"-е",
				},
			},
			{
				gender: Androgynous,
				test: []string{
					"а",
				},
				mods: []string{
					"-ы",
					"-е",
					"-у",
					"-ой",
					"-е",
				},
			},
			{
				gender: Female,
				test: []string{
					"ка",
					"га",
					"ха",
				},
				mods: []string{
					"-и",
					"-е",
					"-у",
					"-ой",
					"-е",
				},
			},
			{
				gender: Female,
				test: []string{
					"ца",
				},
				mods: []string{
					"-ы",
					"-е",
					"-у",
					"-ей",
					"-е",
				},
			},
			{
				gender: Female,
				test: []string{
					"а",
				},
				mods: []string{
					"-ы",
					"-е",
					"-у",
					"-ой",
					"-е",
				},
			},
			{
				gender: Female,
				test: []string{
					"ия",
				},
				mods: []string{
					"-и",
					"-и",
					"-ю",
					"-ей",
					"-и",
				},
			},
			{
				gender: Androgynous,
				test: []string{
					"я",
				},
				mods: []string{
					"-и",
					"-е",
					"-ю",
					"-ей",
					"-е",
				},
			},
			{
				gender: Male,
				test: []string{
					"ий",
				},
				mods: []string{
					"-я",
					"-ю",
					"-я",
					"-ем",
					"-и",
				},
			},
			{
				gender: Male,
				test: []string{
					"ей",
					"й",
				},
				mods: []string{
					"-я",
					"-ю",
					"-я",
					"-ем",
					"-е",
				},
			},
			{
				gender: Male,
				test: []string{
					"ш",
					"ж",
				},
				mods: []string{
					"а",
					"у",
					"а",
					"ем",
					"е",
				},
			},
			{
				gender: Male,
				test: []string{
					"б",
					"в",
					"г",
					"д",
					"ж",
					"з",
					"к",
					"л",
					"м",
					"н",
					"п",
					"р",
					"с",
					"т",
					"ф",
					"х",
					"ц",
					"ч",
				},
				mods: []string{
					"а",
					"у",
					"а",
					"ом",
					"е",
				},
			},
			{
				gender: Androgynous,
				test: []string{
					"ния",
					"рия",
					"вия",
				},
				mods: []string{
					"-и",
					"-и",
					"-ю",
					"-ем",
					"-ем",
				},
			},
		},
	},
	middleName: rulesGroup{
		exceptions: []rule{
			{
				gender: Androgynous,
				test: []string{
					"борух",
				},
				mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
		},
		suffixes: []rule{
			{
				gender: Male,
				test: []string{
					"мич",
					"ьич",
					"кич",
				},
				mods: []string{
					"а",
					"у",
					"а",
					"ом",
					"е",
				},
			},
			{
				gender: Male,
				test: []string{
					"ич",
				},
				mods: []string{
					"а",
					"у",
					"а",
					"ем",
					"е",
				},
			},
			{
				gender: Female,
				test: []string{
					"на",
				},
				mods: []string{
					"-ы",
					"-е",
					"-у",
					"-ой",
					"-е",
				},
			},
		},
	},
	lastName: rulesGroup{
		exceptions: []rule{
			{
				gender: Androgynous,
				test: []string{
					"бонч",
					"абдул",
					"белиц",
					"гасан",
					"дюссар",
					"дюмон",
					"книппер",
					"корвин",
					"ван",
					"шолом",
					"тер",
					"призван",
					"мелик",
					"вар",
					"фон",
				},
				mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
			{
				gender: Androgynous,
				test: []string{
					"дюма",
					"тома",
					"дега",
					"люка",
					"ферма",
					"гамарра",
					"петипа",
					"шандра",
					"скаля",
					"каруана",
				},
				mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
			{
				gender: Androgynous,
				test: []string{
					"гусь",
					"ремень",
					"камень",
					"онук",
					"богода",
					"нечипас",
					"долгопалец",
					"маненок",
					"рева",
					"кива",
				},
				mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
			{
				gender: Male,
				test: []string{
					"вий",
					"сой",
					"цой",
					"хой",
				},
				mods: []string{
					"-я",
					"-ю",
					"-я",
					"-ем",
					"-е",
				},
			},
			{
				gender: Male,
				test: []string{
					"грин",
					"дарвин",
					"регин",
					"цин",
				},
				mods: []string{
					"а",
					"у",
					"а",
					"ом",
					"е",
				},
			},
		},
		suffixes: []rule{
			{
				gender: Female,
				test: []string{
					"б",
					"в",
					"г",
					"д",
					"ж",
					"з",
					"й",
					"к",
					"л",
					"м",
					"н",
					"п",
					"р",
					"с",
					"т",
					"ф",
					"х",
					"ц",
					"ч",
					"ш",
					"щ",
					"ъ",
					"ь",
				},
				mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
			{
				gender: Androgynous,
				test: []string{
					"орота",
				},
				mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
			{
				gender: Female,
				test: []string{
					"ска",
					"цка",
				},
				mods: []string{
					"-ой",
					"-ой",
					"-ую",
					"-ой",
					"-ой",
				},
			},
			{
				gender: Female,
				test: []string{
					"чая",
				},
				mods: []string{
					"--ей",
					"--ей",
					"--ую",
					"--ей",
					"--ей",
				},
			},
			{
				gender: Male,
				test: []string{
					"чий",
				},
				mods: []string{
					"--его",
					"--ему",
					"--его",
					"--им",
					"--ем",
				},
			},
			{
				gender: Female,
				test: []string{
					"цкая",
					"ская",
					"ная",
					"ая",
				},
				mods: []string{
					"--ой",
					"--ой",
					"--ую",
					"--ой",
					"--ой",
				},
			},
			{
				gender: Female,
				test: []string{
					"яя",
				},
				mods: []string{
					"--ей",
					"--ей",
					"--юю",
					"--ей",
					"--ей",
				},
			},
			{
				gender: Male,
				test: []string{
					"иной",
					"уй",
				},
				mods: []string{
					"-я",
					"-ю",
					"-я",
					"-ем",
					"-е",
				},
			},
			{
				gender: Androgynous,
				test: []string{
					"ца",
				},
				mods: []string{
					"-ы",
					"-е",
					"-у",
					"-ей",
					"-е",
				},
			},
			{
				gender: Male,
				test: []string{
					"рих",
				},
				mods: []string{
					"а",
					"у",
					"а",
					"ом",
					"е",
				},
			},
			{
				gender: Androgynous,
				test: []string{
					"ия",
				},
				mods: []string{
					"-и",
					"-и",
					"-ю",
					"-ей",
					"-и",
				},
			},
			{
				gender: Androgynous,
				test: []string{
					"иа",
					"аа",
					"оа",
					"уа",
					"ыа",
					"еа",
					"юа",
					"эа",
				},
				mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
			{
				gender: Androgynous,
				test: []string{
					"о",
					"е",
					"э",
					"и",
					"ы",
					"у",
					"ю",
				},
				mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
			{
				gender: Male,
				test: []string{
					"их",
					"ых",
				},
				mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
			{
				gender: Female,
				test: []string{
					"ова",
					"ева",
					"на",
					"ёва",
				},
				mods: []string{
					"-ой",
					"-ой",
					"-у",
					"-ой",
					"-ой",
				},
			},
			{
				gender: Androgynous,
				test: []string{
					"га",
					"ка",
					"ха",
					"ча",
					"ща",
					"жа",
					"ша",
				},
				mods: []string{
					"-и",
					"-е",
					"-у",
					"-ой",
					"-е",
				},
			},
			{
				gender: Androgynous,
				test: []string{
					"а",
				},
				mods: []string{
					"-ы",
					"-е",
					"-у",
					"-ой",
					"-е",
				},
			},
			{
				gender: Male,
				test: []string{
					"ь",
				},
				mods: []string{
					"-я",
					"-ю",
					"-я",
					"-ем",
					"-е",
				},
			},
			{
				gender: Androgynous,
				test: []string{
					"я",
				},
				mods: []string{
					"-и",
					"-е",
					"-ю",
					"-ей",
					"-е",
				},
			},
			{
				gender: Male,
				test: []string{
					"обей",
				},
				mods: []string{
					"--ья",
					"--ью",
					"--ья",
					"--ьем",
					"--ье",
				},
			},
			{
				gender: Male,
				test: []string{
					"ей",
				},
				mods: []string{
					"-я",
					"-ю",
					"-я",
					"-ем",
					"-е",
				},
			},
			{
				gender: Male,
				test: []string{
					"ян",
					"ан",
					"йн",
				},
				mods: []string{
					"а",
					"у",
					"а",
					"ом",
					"е",
				},
			},
			{
				gender: Male,
				test: []string{
					"ынец",
				},
				mods: []string{
					"--ца",
					"--цу",
					"--ца",
					"--цом",
					"--це",
				},
			},
			{
				gender: Male,
				test: []string{
					"нец",
					"робец",
				},
				mods: []string{
					"--ца",
					"--цу",
					"--ца",
					"--цем",
					"--це",
				},
			},
			{
				gender: Male,
				test: []string{
					"ай",
				},
				mods: []string{
					"-я",
					"-ю",
					"-я",
					"-ем",
					"-е",
				},
			},
			{
				gender: Male,
				test: []string{
					"гой",
					"кой",
				},
				mods: []string{
					"-го",
					"-му",
					"-го",
					"--им",
					"-м",
				},
			},
			{
				gender: Male,
				test: []string{
					"ой",
				},
				mods: []string{
					"-го",
					"-му",
					"-го",
					"--ым",
					"-м",
				},
			},
			{
				gender: Male,
				test: []string{
					"ах",
					"ив",
					"шток",
				},
				mods: []string{
					"а",
					"у",
					"а",
					"ом",
					"е",
				},
			},
			{
				gender: Male,
				test: []string{
					"ший",
					"щий",
					"жий",
					"ний",
				},
				mods: []string{
					"--его",
					"--ему",
					"--его",
					"-м",
					"--ем",
				},
			},
			{
				gender: Male,
				test: []string{
					"ый",
					"кий",
					"хий",
				},
				mods: []string{
					"--ого",
					"--ому",
					"--ого",
					"-м",
					"--ом",
				},
			},
			{
				gender: Male,
				test: []string{
					"ий",
				},
				mods: []string{
					"-я",
					"-ю",
					"-я",
					"-ем",
					"-и",
				},
			},
			{
				gender: Male,
				test: []string{
					"ок",
				},
				mods: []string{
					"--ка",
					"--ку",
					"--ка",
					"--ком",
					"--ке",
				},
			},
			{
				gender: Male,
				test: []string{
					"обец",
					"швец",
					"ьвец",
				},
				mods: []string{
					"а",
					"у",
					"а",
					"ем",
					"е",
				},
			},
			{
				gender: Male,
				test: []string{
					"аец",
					"иец",
					"еец",
				},
				mods: []string{
					"--йца",
					"--йцу",
					"--йца",
					"--йцем",
					"--йце",
				},
			},
			{
				gender: Male,
				test: []string{
					"опец",
				},
				mods: []string{
					"--ца",
					"--цу",
					"--ца",
					"--цем",
					"--це",
				},
			},
			{
				gender: Male,
				test: []string{
					"вец",
					"убец",
					"ырец",
				},
				mods: []string{
					"--ца",
					"--цу",
					"--ца",
					"--цом",
					"--це",
				},
			},
			{
				gender: Male,
				test: []string{
					"ц",
					"ч",
					"ш",
					"щ",
				},
				mods: []string{
					"а",
					"у",
					"а",
					"ем",
					"е",
				},
			},
			{
				gender: Male,
				test: []string{
					"ен",
					"нн",
					"он",
					"ун",
					"б",
					"г",
					"д",
					"ж",
					"з",
					"к",
					"л",
					"м",
					"п",
					"р",
					"с",
					"т",
					"ф",
					"х",
				},
				mods: []string{
					"а",
					"у",
					"а",
					"ом",
					"е",
				},
			},
			{
				gender: Male,
				test: []string{
					"в",
					"н",
				},
				mods: []string{
					"а",
					"у",
					"а",
					"ым",
					"е",
				},
			},
		},
	},
}
