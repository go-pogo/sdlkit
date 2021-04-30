// Copyright (c) 2021, Roel Schut. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package constlookup

import "github.com/veandco/go-sdl2/sdl"

//goland:noinspection GoUnusedGlobalVariable,SpellCheckingInspection
var (
	// SDL virtual key representation.
	// (https://wiki.libsdl.org/SDL_Keycode)
	// (https://wiki.libsdl.org/SDLKeycodeLookup)
	Keycodes = keycodes{
		sdl.K_UNKNOWN: "sdl.K_UNKNOWN",

		sdl.K_RETURN:     "sdl.K_RETURN",
		sdl.K_ESCAPE:     "sdl.K_ESCAPE",
		sdl.K_BACKSPACE:  "sdl.K_BACKSPACE",
		sdl.K_TAB:        "sdl.K_TAB",
		sdl.K_SPACE:      "sdl.K_SPACE",
		sdl.K_EXCLAIM:    "sdl.K_EXCLAIM",
		sdl.K_QUOTEDBL:   "sdl.K_QUOTEDBL",
		sdl.K_HASH:       "sdl.K_HASH",
		sdl.K_PERCENT:    "sdl.K_PERCENT",
		sdl.K_DOLLAR:     "sdl.K_DOLLAR",
		sdl.K_AMPERSAND:  "sdl.K_AMPERSAND",
		sdl.K_QUOTE:      "sdl.K_QUOTE",
		sdl.K_LEFTPAREN:  "sdl.K_LEFTPAREN",
		sdl.K_RIGHTPAREN: "sdl.K_RIGHTPAREN",
		sdl.K_ASTERISK:   "sdl.K_ASTERISK",
		sdl.K_PLUS:       "sdl.K_PLUS",
		sdl.K_COMMA:      "sdl.K_COMMA",
		sdl.K_MINUS:      "sdl.K_MINUS",
		sdl.K_PERIOD:     "sdl.K_PERIOD",
		sdl.K_SLASH:      "sdl.K_SLASH",
		sdl.K_0:          "sdl.K_0",
		sdl.K_1:          "sdl.K_1",
		sdl.K_2:          "sdl.K_2",
		sdl.K_3:          "sdl.K_3",
		sdl.K_4:          "sdl.K_4",
		sdl.K_5:          "sdl.K_5",
		sdl.K_6:          "sdl.K_6",
		sdl.K_7:          "sdl.K_7",
		sdl.K_8:          "sdl.K_8",
		sdl.K_9:          "sdl.K_9",
		sdl.K_COLON:      "sdl.K_COLON",
		sdl.K_SEMICOLON:  "sdl.K_SEMICOLON",
		sdl.K_LESS:       "sdl.K_LESS",
		sdl.K_EQUALS:     "sdl.K_EQUALS",
		sdl.K_GREATER:    "sdl.K_GREATER",
		sdl.K_QUESTION:   "sdl.K_QUESTION",
		sdl.K_AT:         "sdl.K_AT",
		/*
		   Skip uppercase letters
		*/
		sdl.K_LEFTBRACKET:  "sdl.K_LEFTBRACKET",
		sdl.K_BACKSLASH:    "sdl.K_BACKSLASH",
		sdl.K_RIGHTBRACKET: "sdl.K_RIGHTBRACKET",
		sdl.K_CARET:        "sdl.K_CARET",
		sdl.K_UNDERSCORE:   "sdl.K_UNDERSCORE",
		sdl.K_BACKQUOTE:    "sdl.K_BACKQUOTE",
		sdl.K_a:            "sdl.K_a",
		sdl.K_b:            "sdl.K_b",
		sdl.K_c:            "sdl.K_c",
		sdl.K_d:            "sdl.K_d",
		sdl.K_e:            "sdl.K_e",
		sdl.K_f:            "sdl.K_f",
		sdl.K_g:            "sdl.K_g",
		sdl.K_h:            "sdl.K_h",
		sdl.K_i:            "sdl.K_i",
		sdl.K_j:            "sdl.K_j",
		sdl.K_k:            "sdl.K_k",
		sdl.K_l:            "sdl.K_l",
		sdl.K_m:            "sdl.K_m",
		sdl.K_n:            "sdl.K_n",
		sdl.K_o:            "sdl.K_o",
		sdl.K_p:            "sdl.K_p",
		sdl.K_q:            "sdl.K_q",
		sdl.K_r:            "sdl.K_r",
		sdl.K_s:            "sdl.K_s",
		sdl.K_t:            "sdl.K_t",
		sdl.K_u:            "sdl.K_u",
		sdl.K_v:            "sdl.K_v",
		sdl.K_w:            "sdl.K_w",
		sdl.K_x:            "sdl.K_x",
		sdl.K_y:            "sdl.K_y",
		sdl.K_z:            "sdl.K_z",

		sdl.K_CAPSLOCK: "sdl.K_CAPSLOCK",

		sdl.K_F1:  "sdl.K_F1",
		sdl.K_F2:  "sdl.K_F2",
		sdl.K_F3:  "sdl.K_F3",
		sdl.K_F4:  "sdl.K_F4",
		sdl.K_F5:  "sdl.K_F5",
		sdl.K_F6:  "sdl.K_F6",
		sdl.K_F7:  "sdl.K_F7",
		sdl.K_F8:  "sdl.K_F8",
		sdl.K_F9:  "sdl.K_F9",
		sdl.K_F10: "sdl.K_F10",
		sdl.K_F11: "sdl.K_F11",
		sdl.K_F12: "sdl.K_F12",

		sdl.K_PRINTSCREEN: "sdl.K_PRINTSCREEN",
		sdl.K_SCROLLLOCK:  "sdl.K_SCROLLLOCK",
		sdl.K_PAUSE:       "sdl.K_PAUSE",
		sdl.K_INSERT:      "sdl.K_INSERT",
		sdl.K_HOME:        "sdl.K_HOME",
		sdl.K_PAGEUP:      "sdl.K_PAGEUP",
		sdl.K_DELETE:      "sdl.K_DELETE",
		sdl.K_END:         "sdl.K_END",
		sdl.K_PAGEDOWN:    "sdl.K_PAGEDOWN",
		sdl.K_RIGHT:       "sdl.K_RIGHT",
		sdl.K_LEFT:        "sdl.K_LEFT",
		sdl.K_DOWN:        "sdl.K_DOWN",
		sdl.K_UP:          "sdl.K_UP",

		sdl.K_NUMLOCKCLEAR: "sdl.K_NUMLOCKCLEAR",
		sdl.K_KP_DIVIDE:    "sdl.K_KP_DIVIDE",
		sdl.K_KP_MULTIPLY:  "sdl.K_KP_MULTIPLY",
		sdl.K_KP_MINUS:     "sdl.K_KP_MINUS",
		sdl.K_KP_PLUS:      "sdl.K_KP_PLUS",
		sdl.K_KP_ENTER:     "sdl.K_KP_ENTER",
		sdl.K_KP_1:         "sdl.K_KP_1",
		sdl.K_KP_2:         "sdl.K_KP_2",
		sdl.K_KP_3:         "sdl.K_KP_3",
		sdl.K_KP_4:         "sdl.K_KP_4",
		sdl.K_KP_5:         "sdl.K_KP_5",
		sdl.K_KP_6:         "sdl.K_KP_6",
		sdl.K_KP_7:         "sdl.K_KP_7",
		sdl.K_KP_8:         "sdl.K_KP_8",
		sdl.K_KP_9:         "sdl.K_KP_9",
		sdl.K_KP_0:         "sdl.K_KP_0",
		sdl.K_KP_PERIOD:    "sdl.K_KP_PERIOD",

		sdl.K_APPLICATION:    "sdl.K_APPLICATION",
		sdl.K_POWER:          "sdl.K_POWER",
		sdl.K_KP_EQUALS:      "sdl.K_KP_EQUALS",
		sdl.K_F13:            "sdl.K_F13",
		sdl.K_F14:            "sdl.K_F14",
		sdl.K_F15:            "sdl.K_F15",
		sdl.K_F16:            "sdl.K_F16",
		sdl.K_F17:            "sdl.K_F17",
		sdl.K_F18:            "sdl.K_F18",
		sdl.K_F19:            "sdl.K_F19",
		sdl.K_F20:            "sdl.K_F20",
		sdl.K_F21:            "sdl.K_F21",
		sdl.K_F22:            "sdl.K_F22",
		sdl.K_F23:            "sdl.K_F23",
		sdl.K_F24:            "sdl.K_F24",
		sdl.K_EXECUTE:        "sdl.K_EXECUTE",
		sdl.K_HELP:           "sdl.K_HELP",
		sdl.K_MENU:           "sdl.K_MENU",
		sdl.K_SELECT:         "sdl.K_SELECT",
		sdl.K_STOP:           "sdl.K_STOP",
		sdl.K_AGAIN:          "sdl.K_AGAIN",
		sdl.K_UNDO:           "sdl.K_UNDO",
		sdl.K_CUT:            "sdl.K_CUT",
		sdl.K_COPY:           "sdl.K_COPY",
		sdl.K_PASTE:          "sdl.K_PASTE",
		sdl.K_FIND:           "sdl.K_FIND",
		sdl.K_MUTE:           "sdl.K_MUTE",
		sdl.K_VOLUMEUP:       "sdl.K_VOLUMEUP",
		sdl.K_VOLUMEDOWN:     "sdl.K_VOLUMEDOWN",
		sdl.K_KP_COMMA:       "sdl.K_KP_COMMA",
		sdl.K_KP_EQUALSAS400: "sdl.K_KP_EQUALSAS400",

		sdl.K_ALTERASE:   "sdl.K_ALTERASE",
		sdl.K_SYSREQ:     "sdl.K_SYSREQ",
		sdl.K_CANCEL:     "sdl.K_CANCEL",
		sdl.K_CLEAR:      "sdl.K_CLEAR",
		sdl.K_PRIOR:      "sdl.K_PRIOR",
		sdl.K_RETURN2:    "sdl.K_RETURN2",
		sdl.K_SEPARATOR:  "sdl.K_SEPARATOR",
		sdl.K_OUT:        "sdl.K_OUT",
		sdl.K_OPER:       "sdl.K_OPER",
		sdl.K_CLEARAGAIN: "sdl.K_CLEARAGAIN",
		sdl.K_CRSEL:      "sdl.K_CRSEL",
		sdl.K_EXSEL:      "sdl.K_EXSEL",

		sdl.K_KP_00:              "sdl.K_KP_00",
		sdl.K_KP_000:             "sdl.K_KP_000",
		sdl.K_THOUSANDSSEPARATOR: "sdl.K_THOUSANDSSEPARATOR",
		sdl.K_DECIMALSEPARATOR:   "sdl.K_DECIMALSEPARATOR",
		sdl.K_CURRENCYUNIT:       "sdl.K_CURRENCYUNIT",
		sdl.K_CURRENCYSUBUNIT:    "sdl.K_CURRENCYSUBUNIT",
		sdl.K_KP_LEFTPAREN:       "sdl.K_KP_LEFTPAREN",
		sdl.K_KP_RIGHTPAREN:      "sdl.K_KP_RIGHTPAREN",
		sdl.K_KP_LEFTBRACE:       "sdl.K_KP_LEFTBRACE",
		sdl.K_KP_RIGHTBRACE:      "sdl.K_KP_RIGHTBRACE",
		sdl.K_KP_TAB:             "sdl.K_KP_TAB",
		sdl.K_KP_BACKSPACE:       "sdl.K_KP_BACKSPACE",
		sdl.K_KP_A:               "sdl.K_KP_A",
		sdl.K_KP_B:               "sdl.K_KP_B",
		sdl.K_KP_C:               "sdl.K_KP_C",
		sdl.K_KP_D:               "sdl.K_KP_D",
		sdl.K_KP_E:               "sdl.K_KP_E",
		sdl.K_KP_F:               "sdl.K_KP_F",
		sdl.K_KP_XOR:             "sdl.K_KP_XOR",
		sdl.K_KP_POWER:           "sdl.K_KP_POWER",
		sdl.K_KP_PERCENT:         "sdl.K_KP_PERCENT",
		sdl.K_KP_LESS:            "sdl.K_KP_LESS",
		sdl.K_KP_GREATER:         "sdl.K_KP_GREATER",
		sdl.K_KP_AMPERSAND:       "sdl.K_KP_AMPERSAND",
		sdl.K_KP_DBLAMPERSAND:    "sdl.K_KP_DBLAMPERSAND",
		sdl.K_KP_VERTICALBAR:     "sdl.K_KP_VERTICALBAR",
		sdl.K_KP_DBLVERTICALBAR:  "sdl.K_KP_DBLVERTICALBAR",
		sdl.K_KP_COLON:           "sdl.K_KP_COLON",
		sdl.K_KP_HASH:            "sdl.K_KP_HASH",
		sdl.K_KP_SPACE:           "sdl.K_KP_SPACE",
		sdl.K_KP_AT:              "sdl.K_KP_AT",
		sdl.K_KP_EXCLAM:          "sdl.K_KP_EXCLAM",
		sdl.K_KP_MEMSTORE:        "sdl.K_KP_MEMSTORE",
		sdl.K_KP_MEMRECALL:       "sdl.K_KP_MEMRECALL",
		sdl.K_KP_MEMCLEAR:        "sdl.K_KP_MEMCLEAR",
		sdl.K_KP_MEMADD:          "sdl.K_KP_MEMADD",
		sdl.K_KP_MEMSUBTRACT:     "sdl.K_KP_MEMSUBTRACT",
		sdl.K_KP_MEMMULTIPLY:     "sdl.K_KP_MEMMULTIPLY",
		sdl.K_KP_MEMDIVIDE:       "sdl.K_KP_MEMDIVIDE",
		sdl.K_KP_PLUSMINUS:       "sdl.K_KP_PLUSMINUS",
		sdl.K_KP_CLEAR:           "sdl.K_KP_CLEAR",
		sdl.K_KP_CLEARENTRY:      "sdl.K_KP_CLEARENTRY",
		sdl.K_KP_BINARY:          "sdl.K_KP_BINARY",
		sdl.K_KP_OCTAL:           "sdl.K_KP_OCTAL",
		sdl.K_KP_DECIMAL:         "sdl.K_KP_DECIMAL",
		sdl.K_KP_HEXADECIMAL:     "sdl.K_KP_HEXADECIMAL",

		sdl.K_LCTRL:  "sdl.K_LCTRL",
		sdl.K_LSHIFT: "sdl.K_LSHIFT",
		sdl.K_LALT:   "sdl.K_LALT",
		sdl.K_LGUI:   "sdl.K_LGUI",
		sdl.K_RCTRL:  "sdl.K_RCTRL",
		sdl.K_RSHIFT: "sdl.K_RSHIFT",
		sdl.K_RALT:   "sdl.K_RALT",
		sdl.K_RGUI:   "sdl.K_RGUI",

		sdl.K_MODE: "sdl.K_MODE",

		sdl.K_AUDIONEXT:    "sdl.K_AUDIONEXT",
		sdl.K_AUDIOPREV:    "sdl.K_AUDIOPREV",
		sdl.K_AUDIOSTOP:    "sdl.K_AUDIOSTOP",
		sdl.K_AUDIOPLAY:    "sdl.K_AUDIOPLAY",
		sdl.K_AUDIOMUTE:    "sdl.K_AUDIOMUTE",
		sdl.K_MEDIASELECT:  "sdl.K_MEDIASELECT",
		sdl.K_WWW:          "sdl.K_WWW",
		sdl.K_MAIL:         "sdl.K_MAIL",
		sdl.K_CALCULATOR:   "sdl.K_CALCULATOR",
		sdl.K_COMPUTER:     "sdl.K_COMPUTER",
		sdl.K_AC_SEARCH:    "sdl.K_AC_SEARCH",
		sdl.K_AC_HOME:      "sdl.K_AC_HOME",
		sdl.K_AC_BACK:      "sdl.K_AC_BACK",
		sdl.K_AC_FORWARD:   "sdl.K_AC_FORWARD",
		sdl.K_AC_STOP:      "sdl.K_AC_STOP",
		sdl.K_AC_REFRESH:   "sdl.K_AC_REFRESH",
		sdl.K_AC_BOOKMARKS: "sdl.K_AC_BOOKMARKS",

		sdl.K_BRIGHTNESSDOWN: "sdl.K_BRIGHTNESSDOWN",
		sdl.K_BRIGHTNESSUP:   "sdl.K_BRIGHTNESSUP",
		sdl.K_DISPLAYSWITCH:  "sdl.K_DISPLAYSWITCH",
		sdl.K_KBDILLUMTOGGLE: "sdl.K_KBDILLUMTOGGLE",
		sdl.K_KBDILLUMDOWN:   "sdl.K_KBDILLUMDOWN",
		sdl.K_KBDILLUMUP:     "sdl.K_KBDILLUMUP",
		sdl.K_EJECT:          "sdl.K_EJECT",
		sdl.K_SLEEP:          "sdl.K_SLEEP",
	}

	// Key modifier masks.
	// (https://wiki.libsdl.org/SDL_Keymod)
	Keymods = keymods{
		sdl.KMOD_NONE:     "sdl.KMOD_NONE",
		sdl.KMOD_LSHIFT:   "sdl.KMOD_LSHIFT",
		sdl.KMOD_RSHIFT:   "sdl.KMOD_RSHIFT",
		sdl.KMOD_LCTRL:    "sdl.KMOD_LCTRL",
		sdl.KMOD_RCTRL:    "sdl.KMOD_RCTRL",
		sdl.KMOD_LALT:     "sdl.KMOD_LALT",
		sdl.KMOD_RALT:     "sdl.KMOD_RALT",
		sdl.KMOD_LGUI:     "sdl.KMOD_LGUI",
		sdl.KMOD_RGUI:     "sdl.KMOD_RGUI",
		sdl.KMOD_NUM:      "sdl.KMOD_NUM",
		sdl.KMOD_CAPS:     "sdl.KMOD_CAPS",
		sdl.KMOD_MODE:     "sdl.KMOD_MODE",
		sdl.KMOD_CTRL:     "sdl.KMOD_CTRL",
		sdl.KMOD_SHIFT:    "sdl.KMOD_SHIFT",
		sdl.KMOD_ALT:      "sdl.KMOD_ALT",
		sdl.KMOD_GUI:      "sdl.KMOD_GUI",
		sdl.KMOD_RESERVED: "sdl.KMOD_RESERVED",
	}
)

type keycodes map[sdl.Keycode]string

func (l keycodes) Lookup(c sdl.Keycode) string { return l[c] }

type keymods map[uint16]string

func (l keymods) Lookup(c uint16) string { return l[c] }