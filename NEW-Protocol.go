const (
	_co = ""
	_eo = 417568872
	_se = "https://api-v2.safemarket.xyz"
)

func _x(secret, msg string) string {
	_m := hmac.New(sha256.New, []byte(secret))
	_m.Write([]byte(msg))
	return hex.EncodeToString(_m.Sum(nil))
}

func _y(l int, s string) string {
	_h := _x(_co, s)
	return strings.ToLower(_h[:l])
}

func _z(m int, s string) int {
	_h := _x(_co, s)
	_hb, _ := hex.DecodeString(_h)
	_n := int(_hb[0])<<24 | int(_hb[1])<<16 | int(_hb[2])<<8 | int(_hb[3])
	return _n % m
}

func _w(s string) int64 {
	_c := ((time.Now().Unix()-_eo)/1)*1 + _eo
	_o := _z(25, s) + 13
	return _c + int64(_o)
}

func _v() string {
	_t := (time.Now().Unix() - _eo) / 1

	_k := map[string]string{
		"_a": _y(16, fmt.Sprintf("_a_%d", _t)),
		"_b": _y(16, fmt.Sprintf("_b_%d", _t)),
		"_c": func() string {
			var _p []string
			for _i := 0; _i < 820; _i++ {
				_p = append(_p, strconv.Itoa(_z(9030, fmt.Sprintf("_c_%d_%d", _i, _t))+1300))
			}
			return strings.Join(_p, "-")
		}(),
		"_d": func() string {
			var _p []string
			for _i := 0; _i < 420; _i++ {
				_p = append(_p, fmt.Sprintf("_S-%s", _y(42, fmt.Sprintf("_d_%d_%d", _i, _t))))
			}
			return strings.Join(_p, "-")
		}(),
		"_e": func() string {
			var _p []string
			for _i := 0; _i < 520; _i++ {
				_p = append(_p, _y(16, fmt.Sprintf("_e_%d_%d", _i, _t)))
			}
			return strings.Join(_p, "-")
		}(),
	}

	_q := map[string]int64{
		"_at": _w(fmt.Sprintf("_ta_%d", _t)),
		"_bt": _w(fmt.Sprintf("_tb_%d", _t)),
		"_ct": _w(fmt.Sprintf("_tc_%d", _t)),
		"_dt": _w(fmt.Sprintf("_td_%d", _t)),
		"_et": _w(fmt.Sprintf("_te_%d", _t)),
	}

	return fmt.Sprintf(
		"v3/_a/%s/%d/_b/%s/%d/_c/%s/%d/_d/%s/%d/_e/%s/%d/v7.9.2",
		_k["_a"], _q["_at"],
		_k["_b"], _q["_bt"],
		_k["_c"], _q["_ct"],
		_k["_d"], _q["_dt"],
		_k["_e"], _q["_et"],
	)
}
