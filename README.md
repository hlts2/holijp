# holijp
CLI for Japanese holiday

## Installation

```
go get github.com/hlts2/holijp
```

## Usage

```
$ holijp --help
holijp is japanese holiday viewer

Usage:
  holijp [flags]

Flags:
  -d, --day int       set day (default -1)
  -h, --help          help for holijp
      --mode string   set mode(ja, en) (default "ja")
  -m, --month int     set month (default -1)
  -y, --year int      set year (default -1)
```

## Example
## Japanese Mode

List of holidays in 2019 and 11th day in Japanese.
```
$ holijp -y 2019 -d 11
+------+-------+-----+------+--------------+
| YEAR | MONTH | DAY | WEEK |     NAME     |
+------+-------+-----+------+--------------+
| 2019 |     2 |  11 | 月   | 建国記念の日   |
| 2019 |     8 |  11 | 日   | 山の日        |
+------+-------+-----+------+--------------+
```

## English Mode

List of holidays in 2019 in English.

```
$ holijp -y 2019 --mode en
+------+-------+-----+----------+---------------------------+
| YEAR | MONTH | DAY |   WEEK   |           NAME            |
+------+-------+-----+----------+---------------------------+
| 2019 |     1 |   1 | Tuesday  | New Year's Day            |
| 2019 |     1 |  14 | Monday   | Coming of Age Day         |
| 2019 |     2 |  11 | Monday   | National Foundation Day   |
| 2019 |     3 |  21 | Thursday | Vernal Equinox Day        |
| 2019 |     4 |  29 | Monday   | Showa Day                 |
| 2019 |     5 |   3 | Friday   | Constitution Memorial Day |
| 2019 |     5 |   4 | Saturday | Greenery Day              |
| 2019 |     5 |   5 | Sunday   | Children's Day            |
| 2019 |     5 |   6 | Monday   | Holiday in lieu           |
| 2019 |     7 |  15 | Monday   | Marine Day                |
| 2019 |     8 |  11 | Sunday   | Mountain Day              |
| 2019 |     8 |  12 | Monday   | Holiday in lieu           |
| 2019 |     9 |  16 | Monday   | Respect for the Aged Day  |
| 2019 |     9 |  23 | Monday   | Autumnal Equinox Day      |
| 2019 |    10 |  14 | Monday   | Health and Sports Day     |
| 2019 |    11 |   3 | Sunday   | National Culture Day      |
| 2019 |    11 |   4 | Monday   | Holiday in lieu           |
| 2019 |    11 |  23 | Saturday | Labor Thanksgiving Day    |
+------+-------+-----+----------+---------------------------+
```

## License
MIT

## Author
Hiroto Funakoshi (hlts2)
