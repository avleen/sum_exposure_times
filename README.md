# sum_exposure_times

This utility scans a directory tree looking for FITS files.
It then prints the total amount of exposure time in each directory.

This is useful when you have a large number of files and need to know the total
time captured.

## Compatibility

`sum_exposure_times` has been tested on Linux and Windows. It should also work
on OS X.
On Windows the code is flagged by Windows Defender and you need to allow it to
run. Running it inside WSL causes no errors.

## Usage
```
  -cpuprofile string
        write cpu profile to file
  -dir string
        directory to start scanning (default ".")
  -ignore string
        directory tree to ignore
  -threads int
        number of threads (default 24)
```

By default this will scan the current directory, using all of the available CPU
cores. This should maximize disk utilization while it runs.

## Example

```
$ go run ./main.go -dir /mnt/d/Astro -ignore /mnt/d/Astro/Temp
Exptime         Directory (hours:minutes)
-----------------------------------------------------
1:02            /mnt/d/Astro/Andromeda/2021.09.27/M_31/Light
5:13            /mnt/d/Astro/Andromeda/2021.12.01/M_31/Light
1:35            /mnt/d/Astro/Andromeda/2021.12.02/Light
4:22            /mnt/d/Astro/C-2022 E3 (ZTF)/LIGHT
0:00            /mnt/d/Astro/Calibration frames/RedCat 51/Minus 10C/2022.08.13
0:12            /mnt/d/Astro/Calibration frames/RedCat 51/Minus 10C/2023.05.30/L-eXtreme/DARKFLAT
0:13            /mnt/d/Astro/Calibration frames/RedCat 51/Minus 10C/2023.05.30/L-eXtreme/FLAT
8:05            /mnt/d/Astro/California Nebula/2021.11.24/Nina_A914_BH/Light
15:40           /mnt/d/Astro/Christmas Tree Cluster/LIGHT
7:20            /mnt/d/Astro/Cygnus Loop/L-eXtreme/Cygnus Loop Panel 1/LIGHT
8:10            /mnt/d/Astro/Cygnus Loop/L-eXtreme/Cygnus Loop Panel 2/LIGHT
8:25            /mnt/d/Astro/Cygnus Loop/L-eXtreme/Cygnus Loop Panel 3/LIGHT
7:00            /mnt/d/Astro/Cygnus Loop/L-eXtreme/Cygnus Loop Panel 4/LIGHT
5:50            /mnt/d/Astro/Dark Shark Nebula/LIGHT
10:48           /mnt/d/Astro/Dark_Shark_Nebula
15:12           /mnt/d/Astro/Dark_Shark_Nebula/LIGHT
1:24            /mnt/d/Astro/East Veil Nebula/2021.10.24
0:05            /mnt/d/Astro/East Veil Nebula/2021.11.04/Light
1:20            /mnt/d/Astro/East Veil Nebula/2021.11.06/Light
1:55            /mnt/d/Astro/East Veil Nebula/2021.11.16/NGC_6992/Light
1:55            /mnt/d/Astro/East Veil Nebula/2021.11.20
2:05            /mnt/d/Astro/Elephant Trunk Nebula/2021.10.30/IC_1396/Light
0:50            /mnt/d/Astro/Elephant Trunk Nebula/2021.11.02/IC_1396/Light
3:10            /mnt/d/Astro/Elephant Trunk Nebula/2021.11.03/IC_1396/Light
0:50            /mnt/d/Astro/Elephant Trunk Nebula/2021.11.14/IC_1396/Light
8:00            /mnt/d/Astro/Flaming_Star/LIGHT
7:30            /mnt/d/Astro/Fossil_Footprint_Nebula
2:05            /mnt/d/Astro/Ghost_of_Cassiopeia
5:20            /mnt/d/Astro/Heart Nebula/2021.09.28/IC_1805/Light
6:12            /mnt/d/Astro/Horsehead_Nebula
2:00            /mnt/d/Astro/Iris Nebula/2021.09.17
6:15            /mnt/d/Astro/Iris Nebula/LIGHT
8:50            /mnt/d/Astro/Jellyfish_Nebula
0:30            /mnt/d/Astro/Leonard/2021.12.01/Good
0:30            /mnt/d/Astro/Leonard/2021.12.01/Light
9:45            /mnt/d/Astro/Lion Nebula/LIGHT
3:10            /mnt/d/Astro/Monkey_Head_Nebula
20:35           /mnt/d/Astro/NGC 7822/LIGHT
1:06            /mnt/d/Astro/North America Nebula/2021.10.05/NGC_7000/Light
0:21            /mnt/d/Astro/North America Nebula/2021.10.05/NGC_7000/Light old
3:50            /mnt/d/Astro/Orion/2021.11.21/M_42/Light
4:18            /mnt/d/Astro/Orion/2021.12.01/M_42/Light
3:25            /mnt/d/Astro/Pacman
11:20           /mnt/d/Astro/Pinwheel Galaxy/LIGHT
3:45            /mnt/d/Astro/Pleiades/2021.11.01/Light
3:05            /mnt/d/Astro/Pleiades/2021.11.02/M_45/Light
0:33            /mnt/d/Astro/Pleiades/2022.01.30
7:09            /mnt/d/Astro/Pleiades/LIGHT
1:20            /mnt/d/Astro/Rosette Nebula/LIGHT
0:09            /mnt/d/Astro/Rosette/2021.12.05/NGC_2244/Light
8:00            /mnt/d/Astro/SH 2-145/LIGHT
4:15            /mnt/d/Astro/Soul Nebula/LIGHT
6:35            /mnt/d/Astro/Soul_Nebula
2:50            /mnt/d/Astro/Spaghetti Nebula Panel 1/LIGHT
2:10            /mnt/d/Astro/Spaghetti Nebula Panel 2/LIGHT
6:34            /mnt/d/Astro/Triangulum Galaxy/2021.10.24/Good
1:12            /mnt/d/Astro/West Veil Nebula/2021.10.03
2:32            /mnt/d/Astro/West Veil Nebula/2021.10.03/Good
5:22            /mnt/d/Astro/West Veil Nebula/2021.10.04/NGC_6960/Light

Total EXPTIME: 273 hours and 16 minutes
```
