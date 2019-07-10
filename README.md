# Predictor

## Server
### Build
```bash
make server
```

### Config
TBD

### Run
```

```

## Client


## Tests
```bash
make tests
```

## Appendix
### Matrices
#### Indexing
```
[         52, 50, 49         ]
[  6,  5,  4,  3,  2,  1,  0 ]
[ 13, 12, 11, 10,  9,  8,  7 ]
[ 20, 19, 18, 17, 16, 15, 14 ]
[ 27, 26, 25, 24, 23, 22, 21 ]
[ 34, 33, 32, 31, 30, 29, 28 ]
[ 41, 40, 39, 38, 37, 36, 35 ]
[ 48, 47, 46, 45, 44, 43, 42 ]

```
### Cards
```
  0:  1  A♥ | 13: 14  A♣ | 26: 27  A♦ | 39: 40  A♠
  1:  2  2♥ | 14: 15  2♣ | 27: 28  2♦ | 40: 41  2♠
  2:  3  3♥ | 15: 16  3♣ | 28: 29  3♦ | 41: 42  3♠
  3:  4  4♥ | 16: 17  4♣ | 29: 30  4♦ | 42: 43  4♠
  4:  5  5♥ | 17: 18  5♣ | 30: 31  5♦ | 43: 44  5♠
  5:  6  6♥ | 18: 19  6♣ | 31: 32  6♦ | 44: 45  6♠
  6:  7  7♥ | 19: 20  7♣ | 32: 33  7♦ | 45: 46  7♠
  7:  8  8♥ | 20: 21  8♣ | 33: 34  8♦ | 46: 47  8♠
  8:  9  9♥ | 21: 22  9♣ | 34: 35  9♦ | 47: 48  9♠
  9: 10 10♥ | 22: 23 10♣ | 35: 36 10♦ | 48: 49 10♠
 10: 11  J♥ | 23: 24  J♣ | 36: 37  J♦ | 49: 50  J♠
 11: 12  Q♥ | 24: 25  Q♣ | 37: 38  Q♦ | 50: 51  Q♠
 12: 13  K♥ | 25: 26  K♣ | 38: 39  K♦ | 51: 52  K♠
```

### Calendar
```
    | Jan  Feb  Mar  Apr  May  Jun  Jul  Aug  Sep  Oct  Nov  Dec
 ---+-----------------------------------------------------------
  1 |  K♠   J♠   9♠   7♠   5♠   3♠   A♠   Q♦  10♦   8♦   6♦   4♦
  2 |  Q♠  10♠   8♠   6♠   4♠   2♠   K♦   J♦   9♦   7♦   5♦   3♦
  3 |  J♠   9♠   7♠   5♠   3♠   A♠   Q♦  10♦   8♦   6♦   4♦   2♦
  4 | 10♠   8♠   6♠   4♠   2♠   K♦   J♦   9♦   7♦   5♦   3♦   A♦
  5 |  9♠   7♠   5♠   3♠   A♠   Q♦  10♦   8♦   6♦   4♦   2♦   K♣
  6 |  8♠   6♠   4♠   2♠   K♦   J♦   9♦   7♦   5♦   3♦   A♦   Q♣
  7 |  7♠   5♠   3♠   A♠   Q♦  10♦   8♦   6♦   4♦   2♦   K♣   J♣
  8 |  6♠   4♠   2♠   K♦   J♦   9♦   7♦   5♦   3♦   A♦   Q♣  10♣
  9 |  5♠   3♠   A♠   Q♦  10♦   8♦   6♦   4♦   2♦   K♣   J♣   9♣
 10 |  4♠   2♠   K♦   J♦   9♦   7♦   5♦   3♦   A♦   Q♣  10♣   8♣
 11 |  3♠   A♠   Q♦  10♦   8♦   6♦   4♦   2♦   K♣   J♣   9♣   7♣
 12 |  2♠   K♦   J♦   9♦   7♦   5♦   3♦   A♦   Q♣  10♣   8♣   6♣
 13 |  A♠   Q♦  10♦   8♦   6♦   4♦   2♦   K♣   J♣   9♣   7♣   5♣
 14 |  K♦   J♦   9♦   7♦   5♦   3♦   A♦   Q♣  10♣   8♣   6♣   4♣
 15 |  Q♦  10♦   8♦   6♦   4♦   2♦   K♣   J♣   9♣   7♣   5♣   3♣
 16 |  J♦   9♦   7♦   5♦   3♦   A♦   Q♣  10♣   8♣   6♣   4♣   2♣
 17 | 10♦   8♦   6♦   4♦   2♦   K♣   J♣   9♣   7♣   5♣   3♣   A♣
 18 |  9♦   7♦   5♦   3♦   A♦   Q♣  10♣   8♣   6♣   4♣   2♣   K♥
 19 |  8♦   6♦   4♦   2♦   K♣   J♣   9♣   7♣   5♣   3♣   A♣   Q♥
 20 |  7♦   5♦   3♦   A♦   Q♣  10♣   8♣   6♣   4♣   2♣   K♥   J♥
 21 |  6♦   4♦   2♦   K♣   J♣   9♣   7♣   5♣   3♣   A♣   Q♥  10♥
 22 |  5♦   3♦   A♦   Q♣  10♣   8♣   6♣   4♣   2♣   K♥   J♥   9♥
 23 |  4♦   2♦   K♣   J♣   9♣   7♣   5♣   3♣   A♣   Q♥  10♥   8♥
 24 |  3♦   A♦   Q♣  10♣   8♣   6♣   4♣   2♣   K♥   J♥   9♥   7♥
 25 |  2♦   K♣   J♣   9♣   7♣   5♣   3♣   A♣   Q♥  10♥   8♥   6♥
 26 |  A♦   Q♣  10♣   8♣   6♣   4♣   2♣   K♥   J♥   9♥   7♥   5♥
 27 |  K♣   J♣   9♣   7♣   5♣   3♣   A♣   Q♥  10♥   8♥   6♥   4♥
 28 |  Q♣  10♣   8♣   6♣   4♣   2♣   K♥   J♥   9♥   7♥   5♥   3♥
 29 |  J♣   9♣   7♣   5♣   3♣   A♣   Q♥  10♥   8♥   6♥   4♥   2♥
 30 | 10♣        6♣   4♣   2♣   K♥   J♥   9♥   7♥   5♥   3♥   A♥
 31 |  9♣        5♣        A♣       10♥   8♥        4♥        🃏
```

### Planet cycles
```
 01/01-21/02 | 22/02-13/04 | 14/04-04/06 | 05/06-26/07 | 27/07-16/09 | 17/09-07/11 | 08/11-31/12
 02/01-22/02 | 23/02-14/04 | 15/04-05/06 | 06/06-27/07 | 28/07-17/09 | 18/09-08/11 | 09/11-01/01
 03/01-23/02 | 24/02-15/04 | 16/04-06/06 | 07/06-28/07 | 29/07-18/09 | 19/09-09/11 | 10/11-02/01
 04/01-24/02 | 25/02-16/04 | 17/04-07/06 | 08/06-29/07 | 30/07-19/09 | 20/09-10/11 | 11/11-03/01
 05/01-25/02 | 26/02-17/04 | 18/04-08/06 | 09/06-30/07 | 31/07-20/09 | 21/09-11/11 | 12/11-04/01
 06/01-26/02 | 27/02-18/04 | 19/04-09/06 | 10/06-31/07 | 01/08-21/09 | 22/09-12/11 | 13/11-05/01
 07/01-27/02 | 28/02-19/04 | 20/04-10/06 | 11/06-01/08 | 02/08-22/09 | 23/09-13/11 | 14/11-06/01
 08/01-28/02 | 29/02-20/04 | 21/04-11/06 | 12/06-02/08 | 03/08-23/09 | 24/09-14/11 | 15/11-07/01
 09/01-29/02 | 01/03-21/04 | 22/04-12/06 | 13/06-03/08 | 04/08-24/09 | 25/09-15/11 | 16/11-08/01
 10/01-01/03 | 02/03-22/04 | 23/04-13/06 | 14/06-04/08 | 05/08-25/09 | 26/09-16/11 | 17/11-09/01
 11/01-02/03 | 03/03-23/04 | 24/04-14/06 | 15/06-05/08 | 06/08-26/09 | 27/09-17/11 | 18/11-10/01
 12/01-03/03 | 04/03-24/04 | 25/04-15/06 | 16/06-06/08 | 07/08-27/09 | 28/09-18/11 | 19/11-11/01
 13/01-04/03 | 05/03-25/04 | 26/04-16/06 | 17/06-07/08 | 08/08-28/09 | 29/09-19/11 | 20/11-12/01
 14/01-05/03 | 06/03-26/04 | 27/04-17/06 | 18/06-08/08 | 09/08-29/09 | 30/09-20/11 | 21/11-13/01
 15/01-06/03 | 07/03-27/04 | 28/04-18/06 | 19/06-09/08 | 10/08-30/09 | 01/10-21/11 | 22/11-14/01
 16/01-07/03 | 08/03-28/04 | 29/04-19/06 | 20/06-10/08 | 11/08-01/10 | 02/10-22/11 | 23/11-15/01
 17/01-08/03 | 09/03-29/04 | 30/04-20/06 | 21/06-11/08 | 12/08-02/10 | 03/10-23/11 | 24/11-16/01
 18/01-09/03 | 10/03-30/04 | 01/05-21/06 | 22/06-12/08 | 13/08-03/10 | 04/10-24/11 | 25/11-17/01
 19/01-10/03 | 11/03-01/05 | 02/05-22/06 | 23/06-13/08 | 14/08-04/10 | 05/10-25/11 | 26/11-18/01
 20/01-11/03 | 12/03-02/05 | 03/05-23/06 | 24/06-14/08 | 15/08-05/10 | 06/10-26/11 | 27/11-19/01
 21/01-12/03 | 13/03-03/05 | 04/05-24/06 | 25/06-15/08 | 16/08-06/10 | 07/10-27/11 | 28/11-20/01
 22/01-13/03 | 14/03-04/05 | 05/05-25/06 | 26/06-16/08 | 17/08-07/10 | 08/10-28/11 | 29/11-21/01
 23/01-14/03 | 15/03-05/05 | 06/05-26/06 | 27/06-17/08 | 18/08-08/10 | 09/10-29/11 | 30/11-22/01
 24/01-15/03 | 16/03-06/05 | 07/05-27/06 | 28/06-18/08 | 19/08-09/10 | 10/10-30/11 | 01/12-23/01
 25/01-16/03 | 17/03-07/05 | 08/05-28/06 | 29/06-19/08 | 20/08-10/10 | 11/10-01/12 | 02/12-24/01
 26/01-17/03 | 18/03-08/05 | 09/05-29/06 | 30/06-20/08 | 21/08-11/10 | 12/10-02/12 | 03/12-25/01
 27/01-18/03 | 19/03-09/05 | 10/05-30/06 | 01/07-21/08 | 22/08-12/10 | 13/10-03/12 | 04/12-26/01
 28/01-19/03 | 20/03-10/05 | 11/05-01/07 | 02/07-22/08 | 23/08-13/10 | 14/10-04/12 | 05/12-27/01
 29/01-20/03 | 21/03-11/05 | 12/05-02/07 | 03/07-23/08 | 24/08-14/10 | 15/10-05/12 | 06/12-28/01
 30/01-21/03 | 22/03-12/05 | 13/05-03/07 | 04/07-24/08 | 25/08-15/10 | 16/10-06/12 | 07/12-29/01
 31/01-22/03 | 23/03-13/05 | 14/05-04/07 | 05/07-25/08 | 26/08-16/10 | 17/10-07/12 | 08/12-30/01
 01/02-23/03 | 24/03-14/05 | 15/05-05/07 | 06/07-26/08 | 27/08-17/10 | 18/10-08/12 | 09/12-31/01
 02/02-24/03 | 25/03-15/05 | 16/05-06/07 | 07/07-27/08 | 28/08-18/10 | 19/10-09/12 | 10/12-01/02
 03/02-25/03 | 26/03-16/05 | 17/05-07/07 | 08/07-28/08 | 29/08-19/10 | 20/10-10/12 | 11/12-02/02
 04/02-26/03 | 27/03-17/05 | 18/05-08/07 | 09/07-29/08 | 30/08-20/10 | 21/10-11/12 | 12/12-03/02
 05/02-27/03 | 28/03-18/05 | 19/05-09/07 | 10/07-30/08 | 31/08-21/10 | 22/10-12/12 | 13/12-04/02
 06/02-28/03 | 29/03-19/05 | 20/05-10/07 | 11/07-31/08 | 01/09-22/10 | 23/10-13/12 | 14/12-05/02
 07/02-29/03 | 30/03-20/05 | 21/05-11/07 | 12/07-01/09 | 02/09-23/10 | 24/10-14/12 | 15/12-06/02
 08/02-30/03 | 31/03-21/05 | 22/05-12/07 | 13/07-02/09 | 03/09-24/10 | 25/10-15/12 | 16/12-07/02
 09/02-31/03 | 01/04-22/05 | 23/05-13/07 | 14/07-03/09 | 04/09-25/10 | 26/10-16/12 | 17/12-08/02
 10/02-01/04 | 02/04-23/05 | 24/05-14/07 | 15/07-04/09 | 05/09-26/10 | 27/10-17/12 | 18/12-09/02
 11/02-02/04 | 03/04-24/05 | 25/05-15/07 | 16/07-05/09 | 06/09-27/10 | 28/10-18/12 | 19/12-10/02
 12/02-03/04 | 04/04-25/05 | 26/05-16/07 | 17/07-06/09 | 07/09-28/10 | 29/10-19/12 | 20/12-11/02
 13/02-04/04 | 05/04-26/05 | 27/05-17/07 | 18/07-07/09 | 08/09-29/10 | 30/10-20/12 | 21/12-12/02
 14/02-05/04 | 06/04-27/05 | 28/05-18/07 | 19/07-08/09 | 09/09-30/10 | 31/10-21/12 | 22/12-13/02
 15/02-06/04 | 07/04-28/05 | 29/05-19/07 | 20/07-09/09 | 10/09-31/10 | 01/11-22/12 | 23/12-14/02
 16/02-07/04 | 08/04-29/05 | 30/05-20/07 | 21/07-10/09 | 11/09-01/11 | 02/11-23/12 | 24/12-15/02
 17/02-08/04 | 09/04-30/05 | 31/05-21/07 | 22/07-11/09 | 12/09-02/11 | 03/11-24/12 | 25/12-16/02
 18/02-09/04 | 10/04-31/05 | 01/06-22/07 | 23/07-12/09 | 13/09-03/11 | 04/11-25/12 | 26/12-17/02
 19/02-10/04 | 11/04-01/06 | 02/06-23/07 | 24/07-13/09 | 14/09-04/11 | 05/11-26/12 | 27/12-18/02
 20/02-11/04 | 12/04-02/06 | 03/06-24/07 | 25/07-14/09 | 15/09-05/11 | 06/11-27/12 | 28/12-19/02
 21/02-12/04 | 13/04-03/06 | 04/06-25/07 | 26/07-15/09 | 16/09-06/11 | 07/11-28/12 | 29/12-20/02
 22/02-13/04 | 14/04-04/06 | 05/06-26/07 | 27/07-16/09 | 17/09-07/11 | 08/11-29/12 | 30/12-21/02
 23/02-14/04 | 15/04-05/06 | 06/06-27/07 | 28/07-17/09 | 18/09-08/11 | 09/11-30/12 | 31/12-22/02
```
