package net_cat

const (
	GreetingsMsg = "Welcome to TCP-Chat!\n" +
		"         _nnnn_\n" +
		"        dGGGGMMb\n" +
		"       @p~qp~~qMb\n" +
		"       M|@||@) M|\n" +
		"       @,----.JM|\n" +
		"      JS^\\__/  qKL\n" +
		"     dZP        qKRb\n" +
		"    dZP          qKKb\n" +
		"   fZP            SMMb\n" +
		"   HZM            MMMM\n" +
		"   FqM            MMMM\n" +
		" __| \".        |\\dS\"qML\n" +
		" |    `.       | `' \\Zq\n" +
		"_)      \\.___.,|     .'\n" +
		"\\____   )MMMMMP|   .'\n" +
		"     `-'       `--'\n" +
		"[ENTER YOUR NAME]: "

	colorReset = "\033[0m"
	colorGreen = "\033[32m"
	colorRed = "\033[31m"
	colorMagneta = "\033[35m"
	colorGray = "\u001b[47;1m"
	msgPattern = "[%v][%v]:%v"
	TimeDefault = "2006-01-02 15:04:05"
)

// MessageModes
const (
	ModeJoinChat = iota
	ModeSendMessage
	ModeLeftChat
)