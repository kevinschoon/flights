# ✈️  Flights ✈️ 

`flights` is a program that helps you _Find The Best Possible Flight_.


## Usage

```
Usage: flights COMMAND [arg...]

 ✈️  Find The Best Possible Flight ✈️

Commands:
  airports     list all possible airports
  routes       list all possible routes
  furthest     show the furthest distance you can travel
  nearby       show airports nearby
  route        find the best possible routes

```

### airports

List all known airports

#### Example

```
$ flights airports
1Z8:ENIWETOK ISLAND:MARSHALL ISLANDS
AAA:ANAA:FRENCH POLYNESIA
AAE:ANNABA:ALGERIA
AAH:AACHEN:GERMANY
AAL:AALBORG:DENMARK
...
ZVA:MIANDRIVAZO:MADAGASCAR
ZVK:SAVANNAKHET:LAOS
ZWA:ANDAPA:MADAGASCAR
ZYL:SYLHET OSMANI:BANGLADESH
ZZU:MZUZU:MALAWI
```

### routes

List all possible routes

```
$ flights routes
AER:SOCHI:RUSSIA->KZN:KAZAN:RUSSIA
ASF:ASTRAKHAN:RUSSIA->KZN:KAZAN:RUSSIA
ASF:ASTRAKHAN:RUSSIA->MRV:MINERALNYE VODY:RUSSIA
CEK:CHELYABINSK:RUSSIA->KZN:KAZAN:RUSSIA
KZN:KAZAN:RUSSIA->AER:SOCHI:RUSSIA
...
TSV:TOWNSVILLE:AUSTRALIA->ISA:MOUNT ISA:AUSTRALIA
WGA:WAGGA WAGGA:AUSTRALIA->MEL:MELBOURNE:AUSTRALIA
WGA:WAGGA WAGGA:AUSTRALIA->SYD:SYDNEY:AUSTRALIA
FRU:BISHKEK:RUSSIA->OSS:OSH:RUSSIA
OSS:OSH:RUSSIA->FRU:BISHKEK:RUSSIA
```

### furthest

Show the furthest distance you can travel to

```
$ flights furthest ORD
17635.263685:ORD:CHICAGO:USA-->JAD:PERTH:AUSTRALIA
```

### nearby

Show nearby airports

```
$ flights nearby -t 100 ORD
29.587434:ORD:CHICAGO:USA-->DPA:WEST CHICAGO:USA
49.317547:ORD:CHICAGO:USA-->UGN:CHICAGO:USA
24.862137:ORD:CHICAGO:USA-->MDW:CHICAGO:USA
```

### route

Show the quickest possible route

```
$ flights route ORD LIM
6108.125430:ORD:CHICAGO:USA-->PTY:PANAMA CITY:PANAMA-->LIM:LIMA:PERU
```

## TODO

* loadable data structure that contains pricing data
* ...??
