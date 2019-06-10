package wordpress

import "blueprintz/global"

// https://api.wordpress.org/core/stable-check/1.0/
// http://displaywp.com/version/

var Versions = global.Versions{
	"5.3",
	"5.2.1",
	"5.2",
	"5.1.1",
	"5.1",
	"5.0.4",
	"5.0.3",
	"5.0.2",
	"5.0.1",
	"5.0",
	"4.9.10",
	"4.9.9",
	"4.9.8",
	"4.9.7",
	"4.9.6",
	"4.9.5",
	"4.9.4",
	"4.9.3",
	"4.9.2",
	"4.9.1",
	"4.9",
	"4.8.7",
	"4.8.6",
	"4.8.5",
	"4.8.4",
	"4.8.3",
	"4.8.2",
	"4.8.1",
	"4.8",
	"4.7.13",
	"4.7.12",
	"4.7.11",
	"4.7.10",
	"4.7.9",
	"4.7.8",
	"4.7.7",
	"4.7.6",
	"4.7.5",
	"4.7.4",
	"4.7.3",
	"4.7.2",
	"4.7.1",
	"4.7",
	"4.6.12",
	"4.6.11",
	"4.6.10",
	"4.6.9",
	"4.6.8",
	"4.6.7",
	"4.6.6",
	"4.6.5",
	"4.6.4",
	"4.6.3",
	"4.6.2",
	"4.6.1",
	"4.6",
	"4.5.17",
	"4.5.16",
	"4.5.15",
	"4.5.14",
	"4.5.13",
	"4.5.12",
	"4.5.11",
	"4.5.10",
	"4.5.9",
	"4.5.8",
	"4.5.7",
	"4.5.6",
	"4.5.5",
	"4.5.4",
	"4.5.3",
	"4.5.2",
	"4.5.1",
	"4.5",
	"4.4.16",
	"4.4.15",
	"4.4.14",
	"4.4.13",
	"4.4.12",
	"4.4.11",
	"4.4.10",
	"4.4.9",
	"4.4.8",
	"4.4.7",
	"4.4.6",
	"4.4.5",
	"4.4.4",
	"4.4.3",
	"4.4.2",
	"4.4.1",
	"4.4",
	"4.3.19",
	"4.3.18",
	"4.3.17",
	"4.3.16",
	"4.3.15",
	"4.3.14",
	"4.3.13",
	"4.3.12",
	"4.3.11",
	"4.3.10",
	"4.3.9",
	"4.3.8",
	"4.3.7",
	"4.3.6",
	"4.3.5",
	"4.3.4",
	"4.3.3",
	"4.3.2",
	"4.3.1",
	"4.3",
	"4.2.21",
	"4.2.20",
	"4.2.19",
	"4.2.18",
	"4.2.17",
	"4.2.16",
	"4.2.15",
	"4.2.14",
	"4.2.13",
	"4.2.12",
	"4.2.11",
	"4.2.10",
	"4.2.9",
	"4.2.8",
	"4.2.7",
	"4.2.6",
	"4.2.5",
	"4.2.4",
	"4.2.3",
	"4.2.2",
	"4.2.1",
	"4.2",
	"4.1.26",
	"4.1.25",
	"4.1.24",
	"4.1.23",
	"4.1.22",
	"4.1.21",
	"4.1.20",
	"4.1.19",
	"4.1.18",
	"4.1.17",
	"4.1.16",
	"4.1.15",
	"4.1.14",
	"4.1.13",
	"4.1.12",
	"4.1.11",
	"4.1.10",
	"4.1.9",
	"4.1.8",
	"4.1.7",
	"4.1.6",
	"4.1.5",
	"4.1.4",
	"4.1.3",
	"4.1.2",
	"4.1.1",
	"4.1",
	"4.0.24",
	"4.0.23",
	"4.0.22",
	"4.0.21",
	"4.0.20",
	"4.0.19",
	"4.0.18",
	"4.0.17",
	"4.0.16",
	"4.0.15",
	"4.0.14",
	"4.0.13",
	"4.0.12",
	"4.0.11",
	"4.0.10",
	"4.0.9",
	"4.0.8",
	"4.0.7",
	"4.0.6",
	"4.0.5",
	"4.0.4",
	"4.0.3",
	"4.0.2",
	"4.0.1",
	"4.0",
	"3.9.27",
	"3.9.26",
	"3.9.25",
	"3.9.24",
	"3.9.23",
	"3.9.22",
	"3.9.21",
	"3.9.20",
	"3.9.19",
	"3.9.18",
	"3.9.17",
	"3.9.16",
	"3.9.15",
	"3.9.14",
	"3.9.13",
	"3.9.12",
	"3.9.11",
	"3.9.10",
	"3.9.9",
	"3.9.8",
	"3.9.7",
	"3.9.6",
	"3.9.5",
	"3.9.4",
	"3.9.3",
	"3.9.2",
	"3.9.1",
	"3.9",
	"3.8.9",
	"3.8.8",
	"3.8.7",
	"3.8.6",
	"3.8.5",
	"3.8.4",
	"3.8.3",
	"3.8.27",
	"3.8.26",
	"3.8.25",
	"3.8.24",
	"3.8.23",
	"3.8.22",
	"3.8.21",
	"3.8.20",
	"3.8.2",
	"3.8.19",
	"3.8.18",
	"3.8.17",
	"3.8.16",
	"3.8.15",
	"3.8.14",
	"3.8.13",
	"3.8.12",
	"3.8.11",
	"3.8.10",
	"3.8.1",
	"3.8",
	"3.7.29",
	"3.7.28",
	"3.7.27",
	"3.7.26",
	"3.7.25",
	"3.7.24",
	"3.7.23",
	"3.7.22",
	"3.7.21",
	"3.7.20",
	"3.7.19",
	"3.7.18",
	"3.7.17",
	"3.7.16",
	"3.7.15",
	"3.7.14",
	"3.7.13",
	"3.7.12",
	"3.7.11",
	"3.7.10",
	"3.7.9",
	"3.7.8",
	"3.7.7",
	"3.7.6",
	"3.7.5",
	"3.7.4",
	"3.7.3",
	"3.7.2",
	"3.7.1",
	"3.7",
	"3.6.1",
	"3.6",
	"3.5.2",
	"3.5.1",
	"3.5",
	"3.4.2",
	"3.4.1",
	"3.4",
	"3.3.3",
	"3.3.2",
	"3.3.1",
	"3.3",
	"3.2.1",
	"3.2",
	"3.1.4",
	"3.1.3",
	"3.1.2",
	"3.1.1",
	"3.1",
	"3.0.6",
	"3.0.5",
	"3.0.4",
	"3.0.3",
	"3.0.2",
	"3.0.1",
	"3.0",
	"2.9.2",
	"2.9.1",
	"2.9",
	"2.8.6",
	"2.8.5",
	"2.8.4",
	"2.8.3",
	"2.8.2",
	"2.8.1",
	"2.8",
	"2.7.1",
	"2.7",
	"2.6.5",
	"2.6.3",
	"2.6.2",
	"2.6.1",
	"2.6",
	"2.5.1",
	"2.5",
	"2.3.3",
	"2.3.2",
	"2.3.1",
	"2.3",
	"2.2.3",
	"2.2.2",
	"2.2.1",
	"2.2",
	"2.1.3",
	"2.1.2",
	"2.1.1",
	"2.1",
	"2.0.11",
	"2.0.10",
	"2.0.9",
	"2.0.8",
	"2.0.7",
	"2.0.6",
	"2.0.5",
	"2.0.4",
	"2.0.3",
	"2.0.2",
	"2.0.1",
	"2.0",
	"1.5.2",
	"1.5.1.3",
	"1.5.1.2",
	"1.5.1",
	"1.5",
	"1.2.2",
	"1.2.1",
	"1.2",
	"1.0.2",
	"1.0.1",
	"1.0",
	"0.72",
	"0.711",
	"0.70",
}
