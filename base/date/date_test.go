package date

import (
	"syscall/js"
	"testing"

	"github.com/realPy/hogosuru/base/baseobject"
	"github.com/realPy/hogosuru/testingutils"
)

func TestMain(m *testing.M) {
	baseobject.SetSyscall()

	m.Run()

}

func TestNew(t *testing.T) {

	if d, err := New(); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "Date", d.ConstructName_())

	}
	if d, err := New("2015-10-21T03:24:00"); testingutils.AssertErr(t, err) {

		testingutils.AssertExpect(t, "Date", d.ConstructName_())

	}

}

func TestNewFromJSObject(t *testing.T) {

	baseobject.Eval("date=new Date")

	if obj, err := baseobject.Get(js.Global(), "date"); testingutils.AssertErr(t, err) {
		if d, err := NewFromJSObject(obj); testingutils.AssertErr(t, err) {

			testingutils.AssertExpect(t, "Date", d.ConstructName_())

		}
	}

}

func TestGetDate(t *testing.T) {

	if d, err := New("2015-10-21T03:24:00"); testingutils.AssertErr(t, err) {

		if value, err := d.GetDate(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 21, value)
		}

	}

}

func TestSetDate(t *testing.T) {

	if d, err := New("2015-10-21T09:24:00"); testingutils.AssertErr(t, err) {

		if err := d.SetDate(23); testingutils.AssertErr(t, err) {

			gd, _ := d.GetDate()
			testingutils.AssertExpect(t, 23, gd)
		}

	}

}

func TestGetDay(t *testing.T) {

	if d, err := New("2015-10-21T03:24:00"); testingutils.AssertErr(t, err) {

		//Wednesday
		if value, err := d.GetDay(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 3, value)
		}

	}

}

func TestGetFullYear(t *testing.T) {

	if d, err := New("2015-10-21T03:24:00"); testingutils.AssertErr(t, err) {

		if value, err := d.GetFullYear(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 2015, value)
		}

	}

}

func TestSetFullYear(t *testing.T) {

	if d, err := New("2015-10-21T09:24:00"); testingutils.AssertErr(t, err) {

		if err := d.SetFullYear(2021); testingutils.AssertErr(t, err) {

			gf, _ := d.GetFullYear()
			testingutils.AssertExpect(t, 2021, gf)
		}

	}

}
func TestGetHours(t *testing.T) {

	if d, err := New("2015-10-21T09:24:00"); testingutils.AssertErr(t, err) {

		if value, err := d.GetHours(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 9, value)
		}

	}

}

func TestSetHours(t *testing.T) {

	if d, err := New("2015-10-21T09:24:00"); testingutils.AssertErr(t, err) {

		if err := d.SetHours(12); testingutils.AssertErr(t, err) {

			gh, _ := d.GetHours()
			testingutils.AssertExpect(t, 12, gh)
		}

	}

}
func TestGetMilliseconds(t *testing.T) {

	if d, err := New("2015-10-21T09:24:00.100"); testingutils.AssertErr(t, err) {

		if value, err := d.GetMilliseconds(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 100, value)
		}

	}

}

func TestSetMilliseconds(t *testing.T) {

	if d, err := New("2015-10-21T09:24:00.100"); testingutils.AssertErr(t, err) {

		if err := d.SetMilliseconds(200); testingutils.AssertErr(t, err) {

			gm, _ := d.GetMilliseconds()
			testingutils.AssertExpect(t, 200, gm)
		}

	}

}
func TestGetMinutes(t *testing.T) {

	if d, err := New("2015-10-21T09:24:00.100"); testingutils.AssertErr(t, err) {

		if value, err := d.GetMinutes(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 24, value)
		}

	}

}

func TestSetMinutes(t *testing.T) {

	if d, err := New("2015-10-21T09:24:33.100"); testingutils.AssertErr(t, err) {

		if err := d.SetMinutes(48); testingutils.AssertErr(t, err) {

			gm, _ := d.GetMinutes()
			testingutils.AssertExpect(t, 48, gm)
		}

	}

}

func TestGetSeconds(t *testing.T) {

	if d, err := New("2015-10-21T09:24:33.100"); testingutils.AssertErr(t, err) {

		if value, err := d.GetSeconds(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 33, value)
		}

	}

}

func TestSetSeconds(t *testing.T) {

	if d, err := New("2015-10-21T09:24:33.100"); testingutils.AssertErr(t, err) {

		if err := d.SetSeconds(48); testingutils.AssertErr(t, err) {

			gs, _ := d.GetSeconds()
			testingutils.AssertExpect(t, 48, gs)
		}

	}

}

func TestGetTime(t *testing.T) {

	if d, err := New("2015-10-21T09:24:33.100"); testingutils.AssertErr(t, err) {

		if err := d.SetTime(1445412273100 + 10000); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "2015-10-21T07:24:43.100Z", d.ToISOString_())
		}

	}

}

func TestSetTime(t *testing.T) {

	if d, err := New("2015-10-21T09:24:33.100"); testingutils.AssertErr(t, err) {

		if err := d.SetTime(1445412273100 + 10000); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "2015-10-21T07:24:43.100Z", d.ToISOString_())
		}

	}

}

func TestGetTimezoneOffset(t *testing.T) {

	if d, err := New("2015-10-21T09:24:33.100"); testingutils.AssertErr(t, err) {

		if value, err := d.GetTimezoneOffset(); testingutils.AssertErr(t, err) {

			if d2, err := New("2021-10-11T19:24:33.100"); testingutils.AssertErr(t, err) {

				if value2, err := d2.GetTimezoneOffset(); testingutils.AssertErr(t, err) {
					testingutils.AssertExpect(t, true, value2 == value)

				}

			}
		}

	}

}

func TestGetUTCDate(t *testing.T) {

	if d, err := New("August 19, 1975 23:15:30 GMT+11:00"); testingutils.AssertErr(t, err) {

		if value, err := d.GetUTCDate(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 19, value)

		}

	}

	if d, err := New("August 19, 1975 23:15:30 GMT-11:00"); testingutils.AssertErr(t, err) {

		if value, err := d.GetUTCDate(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 20, value)

		}

	}

}

func TestSetUTCDate(t *testing.T) {

	if d, err := New("August 19, 1975 23:15:30 GMT-11:00"); testingutils.AssertErr(t, err) {

		if err := d.SetUTCDate(23); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "1975-08-23T10:15:30.000Z", d.ToISOString_())
		}

	}

}

func TestGetUTCDay(t *testing.T) {

	if d, err := New("August 19, 1975 23:15:30 GMT+11:00"); testingutils.AssertErr(t, err) {

		if value, err := d.GetUTCDay(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 2, value)

		}

	}

	if d, err := New("August 19, 1975 23:15:30 GMT-11:00"); testingutils.AssertErr(t, err) {

		if value, err := d.GetUTCDay(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 3, value)

		}

	}

}

func TestGetUTCFullYear(t *testing.T) {

	if d, err := New("August 19, 1975 23:15:30 GMT+11:00"); testingutils.AssertErr(t, err) {

		if value, err := d.GetUTCFullYear(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 1975, value)

		}

	}

}

func TestSetUTCFullYear(t *testing.T) {

	if d, err := New("August 19, 1975 23:15:30 GMT+11:00"); testingutils.AssertErr(t, err) {

		if err := d.SetUTCFullYear(2021); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "2021-08-19T12:15:30.000Z", d.ToISOString_())
		}

	}

}
func TestGetUTCHours(t *testing.T) {

	if d, err := New("August 19, 1975 23:15:30 GMT+11:00"); testingutils.AssertErr(t, err) {

		if value, err := d.GetUTCHours(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 12, value)

		}

	}

	if d, err := New("August 19, 1975 23:15:30 GMT-11:00"); testingutils.AssertErr(t, err) {

		if value, err := d.GetUTCHours(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 10, value)

		}

	}

}

func TestSetUTCHours(t *testing.T) {

	if d, err := New("August 19, 1975 23:15:30 GMT-11:00"); testingutils.AssertErr(t, err) {

		if err := d.SetUTCHours(20); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "1975-08-20T20:15:30.000Z", d.ToISOString_())
		}

	}

}

func TestGetUTCMilliseconds(t *testing.T) {

	if d, err := New("2018-01-02T03:04:05.678Z"); testingutils.AssertErr(t, err) {

		if value, err := d.GetUTCMilliseconds(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 678, value)

		}

	}

}

func TestSetUTCMilliseconds(t *testing.T) {

	if d, err := New("2018-01-02T03:04:05.678Z"); testingutils.AssertErr(t, err) {

		if err := d.SetUTCMilliseconds(10); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "2018-01-02T03:04:05.010Z", d.ToISOString_())
		}

	}

}

func TestGetUTCMinutes(t *testing.T) {

	if d, err := New("1 January 2000 03:15:30 GMT+07:00"); testingutils.AssertErr(t, err) {

		if value, err := d.GetUTCMinutes(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 15, value)

		}

	}

	if d, err := New("1 January 2000 03:15:30 GMT+03:30"); testingutils.AssertErr(t, err) {

		if value, err := d.GetUTCMinutes(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 45, value)

		}

	}

}

func TestSetUTCMinutes(t *testing.T) {

	if d, err := New("1 January 2000 03:15:30 GMT+03:30"); testingutils.AssertErr(t, err) {

		if err := d.SetUTCMinutes(10); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "1999-12-31T23:10:30.000Z", d.ToISOString_())
		}

	}

}

func TestGetUTCMonth(t *testing.T) {

	if d, err := New("December 31, 1975, 23:15:30 GMT+11:00"); testingutils.AssertErr(t, err) {

		if value, err := d.GetUTCMonth(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 11, value)

		}

	}

	if d, err := New("December 31, 1975, 23:15:30 GMT-11:00"); testingutils.AssertErr(t, err) {

		if value, err := d.GetUTCMonth(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 0, value)

		}

	}

}

func TestSetUTCMonth(t *testing.T) {

	if d, err := New("December 31, 1975, 23:15:30 GMT-11:00"); testingutils.AssertErr(t, err) {

		if err := d.SetUTCMonth(10); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "1976-11-01T10:15:30.000Z", d.ToISOString_())
		}

	}

}

func TestGetUTCSeconds(t *testing.T) {

	if d, err := New("July 20, 1969, 20:18:04 UTC"); testingutils.AssertErr(t, err) {

		if value, err := d.GetUTCSeconds(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, 4, value)

		}

	}

}

func TestSetUTCSeconds(t *testing.T) {

	if d, err := New("July 20, 1969, 20:18:04 UTC"); testingutils.AssertErr(t, err) {

		if err := d.SetUTCSeconds(10); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "1969-07-20T20:18:10.000Z", d.ToISOString_())
		}

	}

}

func TestDateNow(t *testing.T) {

	if d, err := Now(); testingutils.AssertErr(t, err) {
		testingutils.AssertExpect(t, true, d > int64(1635496693781))

	}

}

func TestParse(t *testing.T) {

	if d, err := Parse("01 Jan 1970 00:00:00 GMT"); testingutils.AssertErr(t, err) {
		testingutils.AssertExpect(t, int64(0), d)
	}

	if d, err := Parse("04 Dec 1995 00:12:00 GMT"); testingutils.AssertErr(t, err) {
		testingutils.AssertExpect(t, int64(818035920000), d)
	}

}

func TestToDateString(t *testing.T) {

	if d, err := New(1993, 6, 28, 14, 39, 7); testingutils.AssertErr(t, err) {

		if s, err := d.ToDateString(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "Wed Jul 28 1993", s)
		}

	}

}

func TestISOString(t *testing.T) {

	if timestamp, err := UTC(1993, 6, 28, 14, 39, 7); testingutils.AssertErr(t, err) {

		if d, err := New(timestamp); testingutils.AssertErr(t, err) {

			if s, err := d.ToISOString(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "1993-07-28T14:39:07.000Z", s)
			}

		}
	}

}

func TestToJSON(t *testing.T) {

	if timestamp, err := UTC(1993, 6, 28, 14, 39, 7); testingutils.AssertErr(t, err) {

		if d, err := New(timestamp); testingutils.AssertErr(t, err) {

			if s, err := d.ToJSON(); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "1993-07-28T14:39:07.000Z", s)
			}

		}
	}

}

func TestToLocaleDateString(t *testing.T) {

	if timestamp, err := UTC(1993, 6, 28, 14, 39, 7); testingutils.AssertErr(t, err) {

		if d, err := New(timestamp); testingutils.AssertErr(t, err) {

			if s, err := d.ToLocaleDateString("fr-FR", map[string]interface{}{"weekday": "long", "year": "numeric", "month": "long", "day": "numeric"}); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "mercredi 28 juillet 1993", s)
			}

		}
	}

}

func TestToLocaleString(t *testing.T) {

	if timestamp, err := UTC(1993, 6, 28, 14, 39, 7); testingutils.AssertErr(t, err) {

		if d, err := New(timestamp); testingutils.AssertErr(t, err) {

			if s, err := d.ToLocaleString("fr-FR", map[string]interface{}{"timeZone": "UTC"}); testingutils.AssertErr(t, err) {
				testingutils.AssertExpect(t, "28/07/1993 14:39:07", s)
			}

		}
	}

}

/*
func TestToLocaleTimeString(t *testing.T) {

	if d, err := New("August 19, 1975 23:15:30 GMT+00:00"); testingutils.AssertErr(t, err) {

		if s, err := d.ToLocaleTimeString("en-US"); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "12:15:30 AM", s)
		}

	}

}

func TestToUTCString(t *testing.T) {

	if d, err := New("August 19, 1975 23:15:30"); testingutils.AssertErr(t, err) {

		if s, err := d.ToUTCString(); testingutils.AssertErr(t, err) {
			testingutils.AssertExpect(t, "Tue, 19 Aug 1975 22:15:30 GMT", s)
		}

	}

}*/
