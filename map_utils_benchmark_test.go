package map_utils

import (
	"sort"
	"testing"
)

var stringMap10 = map[string]interface{}{
	"a94a8fe5ccb19ba61c4c0873d391e987982fbbd3": map[string]interface{}{"foo": "bar", "test": 567.87},
	"0beec7b5ea3f0fdbc95d0dd47f3c5bc275da8a33": map[string]interface{}{"foo": []string{"bar", "baz", "qux"}},
	"62cdb7020ff920e5aa642c3d4066950dd1f01f4d": []int{1, 2, 34, 6, 66787, 90909, 34342, 4},
	"bbe960a25ea311d21d40669e93df2003ba9b90a2": []map[string]interface{}{{"foo": "bar", "baz": 123}},
	"b54ba7f5621240d403f06815f7246006ef8c7d43": map[string]interface{}{"foo": "bar", "test": 567.87},
	"3da541559918a808c2402bba5012f6c60b27661c": map[string]interface{}{"foo": []string{"bar", "baz", "qux"}},
	"dc724af18fbdd4e59189f5fe768a5f8311527050": []int{1, 2, 34, 6, 66787, 90909, 34342, 4},
	"35eebf6635d6babe24b031dcaf4f4a55ebe8da9c": []map[string]interface{}{{"foo": "bar", "baz": 123}},
	"675b25d054d129341c911c93a4e866e03feb0196": map[string]interface{}{"foo": "bar"},
	"438826866698ca0e9ccb304bd87238740c532198": map[string]interface{}{"baz": []int{345, 6, 62343}},
}

var stringMap100 = map[string]interface{}{
	"r5MDoC3aflcLl24CB5pDx5Q6KL6xPhlXgiKCC4G8": "r5MDoC3aflcLl24CB5pDx5Q6KL6xPhlXgiKCC4G8",
	"TQCiUEkTgcpt4DKEnU76ryymbpmJfBArGMXYiPRi": "TQCiUEkTgcpt4DKEnU76ryymbpmJfBArGMXYiPRi",
	"thhu9FMJad1bVLhdie5CqhbyrDVjYSq4sAXKGCm5": "thhu9FMJad1bVLhdie5CqhbyrDVjYSq4sAXKGCm5",
	"hPG9nCDpDaJVeKioww09EmPwIfG8RBFxTybxfgl4": "hPG9nCDpDaJVeKioww09EmPwIfG8RBFxTybxfgl4",
	"bPIuQsITw5SDr9IepP6wygHGXLd5JdrEsVFJWjgL": "bPIuQsITw5SDr9IepP6wygHGXLd5JdrEsVFJWjgL",
	"G4LRhso5s5C5YmUjm2ofsco7ZFQ2DjccTJbA71Z6": "G4LRhso5s5C5YmUjm2ofsco7ZFQ2DjccTJbA71Z6",
	"fMeBn8alCJLrk3xE77w52fRRr4CwG8ZPSvpM1Ucr": "fMeBn8alCJLrk3xE77w52fRRr4CwG8ZPSvpM1Ucr",
	"yJlV3MKc0oieMBx2i7fQtoAcDFKhY3eFH7TIFRFx": "yJlV3MKc0oieMBx2i7fQtoAcDFKhY3eFH7TIFRFx",
	"W27AiVCkVJI5g6dZ3WwZSuWyWZLg04quAeDBMqqO": "W27AiVCkVJI5g6dZ3WwZSuWyWZLg04quAeDBMqqO",
	"fU8OrHMn44HWZXstiYXHcq0DwvEj9kuCIPBiBNSg": "fU8OrHMn44HWZXstiYXHcq0DwvEj9kuCIPBiBNSg",
	"Z9GKrdXm7Ln9IbbwUdxboMQOf08EMVq2Wy3KxNPf": "Z9GKrdXm7Ln9IbbwUdxboMQOf08EMVq2Wy3KxNPf",
	"Qp36mmlJbAqomXbbx3YT6MuIWdTReV5E32XpOCwn": "Qp36mmlJbAqomXbbx3YT6MuIWdTReV5E32XpOCwn",
	"DtHihwiVrreblAhzVrnZNpi8aDb8rltKPNkKpv2a": "DtHihwiVrreblAhzVrnZNpi8aDb8rltKPNkKpv2a",
	"VEYQGV8VdBYutHRh76Wgy6YM4Br8ZWw7ak2BsGOa": "VEYQGV8VdBYutHRh76Wgy6YM4Br8ZWw7ak2BsGOa",
	"LbpQQsB1WDK5Gy6qo0rGy0dbm8lIjDHihVkpgwkh": "LbpQQsB1WDK5Gy6qo0rGy0dbm8lIjDHihVkpgwkh",
	"VgN8cdU5QSTXE1XmvoBv0kBL6c1sKS59mPftVWS4": "VgN8cdU5QSTXE1XmvoBv0kBL6c1sKS59mPftVWS4",
	"9ui8OYIcGfJD1tLIF2Wf5E766mf5OB2xvj1iD0aT": "9ui8OYIcGfJD1tLIF2Wf5E766mf5OB2xvj1iD0aT",
	"TDc2yyAKoV42hFW6yP0bXqja04LFvVrNgrobayUj": "TDc2yyAKoV42hFW6yP0bXqja04LFvVrNgrobayUj",
	"peVE6k51BOaGCgiz3wyFtzjSkp63PKgkvSAV63Hf": "peVE6k51BOaGCgiz3wyFtzjSkp63PKgkvSAV63Hf",
	"AHnmE2Nq3OWJS3ZUdWuKvHqJA6bNcWhXmMVvFeZA": "AHnmE2Nq3OWJS3ZUdWuKvHqJA6bNcWhXmMVvFeZA",
	"tNFNo5DkpAbEtYG5NQ8al5lT5gkdftsWmPwUe96q": "tNFNo5DkpAbEtYG5NQ8al5lT5gkdftsWmPwUe96q",
	"4y2qV4kLAz2YCeWnUYsKlaBgRvm73H8AtfkC6WYs": "4y2qV4kLAz2YCeWnUYsKlaBgRvm73H8AtfkC6WYs",
	"845PYE0uTroOQZuxCH1PxunTbRnq5I2HqroZm96s": "845PYE0uTroOQZuxCH1PxunTbRnq5I2HqroZm96s",
	"A6L8ZBs8Lk777pzFKtqZ1HYFPqcNNrMdC3UpJ0Vp": "A6L8ZBs8Lk777pzFKtqZ1HYFPqcNNrMdC3UpJ0Vp",
	"nqkjU2LtKFJ1f8ovlLtgAJbvnYPJ3MNhfX88En1Y": "nqkjU2LtKFJ1f8ovlLtgAJbvnYPJ3MNhfX88En1Y",
	"wqVf5qt5druXyj4Gq5fYrN8jIdW86bgex2w0k8ei": "wqVf5qt5druXyj4Gq5fYrN8jIdW86bgex2w0k8ei",
	"tpt6wbMMTr3nRGrQ3bBakpN96cuZeB2GyHpVOBO6": "tpt6wbMMTr3nRGrQ3bBakpN96cuZeB2GyHpVOBO6",
	"mkZq4lFkL5xGBhtGiCUtbyK7wjDPDJ72RnRvCBbs": "mkZq4lFkL5xGBhtGiCUtbyK7wjDPDJ72RnRvCBbs",
	"b9jgXe30vqtSCQko0IZtHRcU0GrMdsDYEXdDDfhE": "b9jgXe30vqtSCQko0IZtHRcU0GrMdsDYEXdDDfhE",
	"0uwt6w77PUKSD9hU61fXBMY8JLmCYs5pf4Z4XTjl": "0uwt6w77PUKSD9hU61fXBMY8JLmCYs5pf4Z4XTjl",
	"jsVpt849R7JIGFzEepoaevwXCuFthEr5mMd4bu6l": "jsVpt849R7JIGFzEepoaevwXCuFthEr5mMd4bu6l",
	"ueEv54u2ye18D0MJIJfAlW410Odv4GYsbOaQyxZ5": "ueEv54u2ye18D0MJIJfAlW410Odv4GYsbOaQyxZ5",
	"iWsHgOhtqNv2RhBNX98DFvzRbfRmcfsgLkT96LTz": "iWsHgOhtqNv2RhBNX98DFvzRbfRmcfsgLkT96LTz",
	"RZhVnSJhSnEBg9IMRbVKYl1wBpinOYMBn7ictEuN": "RZhVnSJhSnEBg9IMRbVKYl1wBpinOYMBn7ictEuN",
	"N579OXpwL4zkXEuBavRWKZC0c6eMXPxTo4uODjiv": "N579OXpwL4zkXEuBavRWKZC0c6eMXPxTo4uODjiv",
	"tbpIDmtxZv7Qo5e2PH3pLEWl3AOD7ObCAeLRA3na": "tbpIDmtxZv7Qo5e2PH3pLEWl3AOD7ObCAeLRA3na",
	"s4yz2FzrIv9HUWETOEYAKcOBgu92BgWHjxBHXJOj": "s4yz2FzrIv9HUWETOEYAKcOBgu92BgWHjxBHXJOj",
	"lturq07qhqpvSuWeDPx5d7BQAn211drBD2cT8tMJ": "lturq07qhqpvSuWeDPx5d7BQAn211drBD2cT8tMJ",
	"N77KnvpFxDn5QO87tRMWvNMbjU9VNle1zz0Y9Iv6": "N77KnvpFxDn5QO87tRMWvNMbjU9VNle1zz0Y9Iv6",
	"0oyq6GPF14bHdaLi3pV221J74ScVxklDvwRQgm7o": "0oyq6GPF14bHdaLi3pV221J74ScVxklDvwRQgm7o",
	"BCy70XSzuEfeP9HVxDi3CFoHcWBUuAW3HwZSSPVI": "BCy70XSzuEfeP9HVxDi3CFoHcWBUuAW3HwZSSPVI",
	"L3Tvt3rEIJln6UktGe8XLyd4mNsGVnoMQ4ZBNSHb": "L3Tvt3rEIJln6UktGe8XLyd4mNsGVnoMQ4ZBNSHb",
	"pY0GlOsPTj4uiZ5qh4XDcefJTmq5E1R9n0siJwAm": "pY0GlOsPTj4uiZ5qh4XDcefJTmq5E1R9n0siJwAm",
	"LW3EbycE0ThBlEjNv2vpdkkVKUvsnnXgAkfu7A0F": "LW3EbycE0ThBlEjNv2vpdkkVKUvsnnXgAkfu7A0F",
	"ABF7t0QfueWxe3B69bHtQYhV8tYYOQ9JMeM0V3Hn": "ABF7t0QfueWxe3B69bHtQYhV8tYYOQ9JMeM0V3Hn",
	"jfzsrU0RhFzAdzjb67lfXy9ntO331WX1r7Yyotj1": "jfzsrU0RhFzAdzjb67lfXy9ntO331WX1r7Yyotj1",
	"JnGS9PQaXSEGlFpugNuMMQCEPSX9TyKQSUVFgT0V": "JnGS9PQaXSEGlFpugNuMMQCEPSX9TyKQSUVFgT0V",
	"e2aUdfTWmUNs8jE6wHuyd6u5LCflPEeuFTGuGkwH": "e2aUdfTWmUNs8jE6wHuyd6u5LCflPEeuFTGuGkwH",
	"O0oKK7M2031XCDc8JFSGmO0be73wvhQEXVydEytY": "O0oKK7M2031XCDc8JFSGmO0be73wvhQEXVydEytY",
	"468O5MLgqLEJozhaboEuwno25fg44V7jxMDAHw65": "468O5MLgqLEJozhaboEuwno25fg44V7jxMDAHw65",
	"VG1VCRyaUkOCmUN7BP0rkv3dUn33cFTa9fn3IW9D": "VG1VCRyaUkOCmUN7BP0rkv3dUn33cFTa9fn3IW9D",
	"EXnlRG3nzSxofStgfSv9Ba3URnYemnDsKezODxQe": "EXnlRG3nzSxofStgfSv9Ba3URnYemnDsKezODxQe",
	"oPnI0yge1BalnBC25KqVYMTLyZ0S8oUK05TDYBmw": "oPnI0yge1BalnBC25KqVYMTLyZ0S8oUK05TDYBmw",
	"ff90HvWxa0wQyKmfDwRcWH7KEfqU61AyHfhzocHZ": "ff90HvWxa0wQyKmfDwRcWH7KEfqU61AyHfhzocHZ",
	"gqlJ85dGhhub27evpJb8YzNsPPZztHtVV13lUl89": "gqlJ85dGhhub27evpJb8YzNsPPZztHtVV13lUl89",
	"ECbzZqfjjH0abnSoh4TkN5cXD6KP5xiYK1HnOFQT": "ECbzZqfjjH0abnSoh4TkN5cXD6KP5xiYK1HnOFQT",
	"McgbOwdODWtItjCGNH5i1Gt6s2PYqAmLUOw1Nil2": "McgbOwdODWtItjCGNH5i1Gt6s2PYqAmLUOw1Nil2",
	"pUbUdSSTGVaZkJRecMpT31RP5idoKqZtSxpiBgLY": "pUbUdSSTGVaZkJRecMpT31RP5idoKqZtSxpiBgLY",
	"6NY3ranAYuIakSe8Tpv3UQHkQT1cLD0oAAfmAvQ4": "6NY3ranAYuIakSe8Tpv3UQHkQT1cLD0oAAfmAvQ4",
	"NKjtFACmxmHxHXESMvw3AcF4QfteGIPigO29GWIa": "NKjtFACmxmHxHXESMvw3AcF4QfteGIPigO29GWIa",
	"LhF8u0sSsIWdV8zlBiQMlVH6hDq4IHbTmApkP4Vy": "LhF8u0sSsIWdV8zlBiQMlVH6hDq4IHbTmApkP4Vy",
	"XdtaHksSG8P0iz0p9NJJTaNuXggelE6jNGTpHRY3": "XdtaHksSG8P0iz0p9NJJTaNuXggelE6jNGTpHRY3",
	"qYlsoWoJ1faWBOLed6SnHHHutKUrldzc3QntqRp2": "qYlsoWoJ1faWBOLed6SnHHHutKUrldzc3QntqRp2",
	"IZ9EXI6ioU108jWVGjve2styprMl8zdDVg0BPoZL": "IZ9EXI6ioU108jWVGjve2styprMl8zdDVg0BPoZL",
	"4tiWZ2IzmLRqk7dAiLQiLfs5erEhN9PFa6Mdnyjq": "4tiWZ2IzmLRqk7dAiLQiLfs5erEhN9PFa6Mdnyjq",
	"n6rA1AzCeSerACtjOtwMWDv2LvG31ch9dsuMa8GN": "n6rA1AzCeSerACtjOtwMWDv2LvG31ch9dsuMa8GN",
	"kXMbRIOTSbcNliE1Rb0n6SE5PBlEWuohF4LTSynS": "kXMbRIOTSbcNliE1Rb0n6SE5PBlEWuohF4LTSynS",
	"vhyFxJTmsEJnT0TP1UZUt7CDeAnuPrJrNeYBqDgS": "vhyFxJTmsEJnT0TP1UZUt7CDeAnuPrJrNeYBqDgS",
	"0pe50t8AHl4fcTJxf38hMlsfTKuvYcTmwEl52BoA": "0pe50t8AHl4fcTJxf38hMlsfTKuvYcTmwEl52BoA",
	"UYoAYj8vCterCUVUsgW1x0UILBYVYQvpSGaacDvd": "UYoAYj8vCterCUVUsgW1x0UILBYVYQvpSGaacDvd",
	"Gl4KX6h6wdrPBiJuZIHl2eRPwA71vtELlWcSQSVV": "Gl4KX6h6wdrPBiJuZIHl2eRPwA71vtELlWcSQSVV",
	"mJVLmiJgKpk00CkG30xA4rFFD3DgwOB6FGzlnM5v": "mJVLmiJgKpk00CkG30xA4rFFD3DgwOB6FGzlnM5v",
	"d9fwBkdqJmW5NaVZcrg537givVq0IOQkWp8G15T7": "d9fwBkdqJmW5NaVZcrg537givVq0IOQkWp8G15T7",
	"IZ513eMksrroYHNhwPPvBTeRa00EM4epj0K3MLi6": "IZ513eMksrroYHNhwPPvBTeRa00EM4epj0K3MLi6",
	"zx7keLeChFAHtf0AuoRmNBbEclCArriFz4WzUSdR": "zx7keLeChFAHtf0AuoRmNBbEclCArriFz4WzUSdR",
	"E5NSQFljgHxlQSfi2h3C0SIQesYV3ZZFnoKfFxAF": "E5NSQFljgHxlQSfi2h3C0SIQesYV3ZZFnoKfFxAF",
	"1hqfPwp6VC98GSqt75cnUt2k4FhFdNjsgHcvwrbw": "1hqfPwp6VC98GSqt75cnUt2k4FhFdNjsgHcvwrbw",
	"CaF3Dx3CsdRCBS92unISrQrd0EZEOlowYy0OZEjd": "CaF3Dx3CsdRCBS92unISrQrd0EZEOlowYy0OZEjd",
	"2rpvN4yKCOZD9aBMeGxxKNQ9JC7I86VifqfPlezs": "2rpvN4yKCOZD9aBMeGxxKNQ9JC7I86VifqfPlezs",
	"OjRywidJa85ay9evxCS1WrJDSVmAfJdqpyMB96Gk": "OjRywidJa85ay9evxCS1WrJDSVmAfJdqpyMB96Gk",
	"myzh9f6ymb5MMJxPZMfZwnX7i6dUXtB15rnus2aH": "myzh9f6ymb5MMJxPZMfZwnX7i6dUXtB15rnus2aH",
	"05AWL0HYppTMER5lWeUJQDZneu7S9nOT3AURhOvC": "05AWL0HYppTMER5lWeUJQDZneu7S9nOT3AURhOvC",
	"uglBwSyUlxYlEZA7wbHn9KsYCZm4VPdVisg7oilJ": "uglBwSyUlxYlEZA7wbHn9KsYCZm4VPdVisg7oilJ",
	"6PNMm7pNBwX0crWSLm3MWOnXG2XKb81ESBhaUanr": "6PNMm7pNBwX0crWSLm3MWOnXG2XKb81ESBhaUanr",
	"H2Sfwgkf2R069n2HaWjBFscb8WkjxrEiRGzivlYB": "H2Sfwgkf2R069n2HaWjBFscb8WkjxrEiRGzivlYB",
	"o7dVXwdre6Ak6zvwgLtMSKsEo5s98UuqqaT0Mepm": "o7dVXwdre6Ak6zvwgLtMSKsEo5s98UuqqaT0Mepm",
	"Sc8FNOraTQ0pbVl1RgG6XVN0y9sWRg2vAK4GV4r0": "Sc8FNOraTQ0pbVl1RgG6XVN0y9sWRg2vAK4GV4r0",
	"PpNMzkGRgxY3qRZ70yNoL9c35NYyMnE3pymNnXd6": "PpNMzkGRgxY3qRZ70yNoL9c35NYyMnE3pymNnXd6",
	"ZqVQdFZmVcAY2JcFSawaXCGynk9rJonElP8yh60w": "ZqVQdFZmVcAY2JcFSawaXCGynk9rJonElP8yh60w",
	"0WZSmfGHFF8laB38fVxTQ4Otvh5qzb1Kx23o4nvq": "0WZSmfGHFF8laB38fVxTQ4Otvh5qzb1Kx23o4nvq",
	"wuvqilkbV2l58wpvyVTA9t7SxQfCYYlWYzM0WCx5": "wuvqilkbV2l58wpvyVTA9t7SxQfCYYlWYzM0WCx5",
	"dv4vZayAcCLNRZwMdIwfqzrIkqnhT3FCp6doRVJr": "dv4vZayAcCLNRZwMdIwfqzrIkqnhT3FCp6doRVJr",
	"Lb3aTyv5GFq2OTAuE2XFgR7v5WpWLmcOTOU3SUGx": "Lb3aTyv5GFq2OTAuE2XFgR7v5WpWLmcOTOU3SUGx",
	"PtH5vC1tExZ4FAsY5QNEU2YFeobPloi9R85X9RLx": "PtH5vC1tExZ4FAsY5QNEU2YFeobPloi9R85X9RLx",
	"RVWF2PkpUO6CwBDRTmbYOiN5dTF78sZXTFq8W9FT": "RVWF2PkpUO6CwBDRTmbYOiN5dTF78sZXTFq8W9FT",
	"Mp8E1whSLUoBh550QtHT8HKRRymIt3EcgZK6maSJ": "Mp8E1whSLUoBh550QtHT8HKRRymIt3EcgZK6maSJ",
	"fFvy6Q5DhoJF9cS1YPdYheqZplx7B55O2Yx779YA": "fFvy6Q5DhoJF9cS1YPdYheqZplx7B55O2Yx779YA",
	"NHWIha1IgogoXM9qmZdObUsz1mC4G8srbCpYkkPP": "NHWIha1IgogoXM9qmZdObUsz1mC4G8srbCpYkkPP",
	"Jyut64qjFjrXu99srGDxl5A9cCIrBAlTC68j62Cn": "Jyut64qjFjrXu99srGDxl5A9cCIrBAlTC68j62Cn",
	"50waEf4MSraFj9bbrXDO2nDNQ0VBB06ATzOgOIbs": "50waEf4MSraFj9bbrXDO2nDNQ0VBB06ATzOgOIbs",
}

func stringKeys(dict map[string]interface{}) (keys []string) {
	for key := range dict {
		keys = append(keys, key)
	}
	return
}

func benchmark_stringKeys(dict map[string]interface{}, b *testing.B) {
	for n := 0; n < b.N; n++ {
		keys := stringKeys(dict)
		for range keys {
		}
	}
}

func benchmark_Keys(dict map[string]interface{}, b *testing.B) {
	for n := 0; n < b.N; n++ {
		var keys []string
		Keys(dict, &keys)
		for range keys {
		}
	}
}

func Benchmark_Keys10(b *testing.B)  { benchmark_Keys(stringMap10, b) }
func Benchmark_Keys100(b *testing.B) { benchmark_Keys(stringMap100, b) }

func Benchmark_stringKeys10(b *testing.B)  { benchmark_stringKeys(stringMap10, b) }
func Benchmark_stringKeys100(b *testing.B) { benchmark_stringKeys(stringMap100, b) }

func sortedStringKeys(dict map[string]interface{}) (keys []string) {
	for key := range dict {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return
}

func benchmark_sortedStringKeys(dict map[string]interface{}, b *testing.B) {
	for n := 0; n < b.N; n++ {
		keys := sortedStringKeys(dict)
		for range keys {
		}
	}
}

func benchmark_SortedKeys(dict map[string]interface{}, b *testing.B) {
	for n := 0; n < b.N; n++ {
		var keys []string
		SortedKeys(dict, &keys)
		for range keys {
		}
	}
}

func benchmark_Keys_sortAfter(dict map[string]interface{}, b *testing.B) {
	for n := 0; n < b.N; n++ {
		var keys []string
		Keys(dict, &keys)
		sort.Strings(keys)
		for range keys {
		}
	}
}

func Benchmark_SortedKeys10(b *testing.B)  { benchmark_SortedKeys(stringMap10, b) }
func Benchmark_SortedKeys100(b *testing.B) { benchmark_SortedKeys(stringMap100, b) }

func Benchmark_Keys_sortAfter10(b *testing.B)  { benchmark_SortedKeys(stringMap10, b) }
func Benchmark_Keys_sortAfter100(b *testing.B) { benchmark_SortedKeys(stringMap100, b) }

func Benchmark_sortedStringKeys10(b *testing.B)  { benchmark_sortedStringKeys(stringMap10, b) }
func Benchmark_sortedStringKeys100(b *testing.B) { benchmark_sortedStringKeys(stringMap100, b) }
