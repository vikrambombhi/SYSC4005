# How to run
If you are using windows run `./simulation_windows`
if you are using OSX run `./simulation_mac`
if you are using Linux run `./simulation_linux`


To specify the directory containing data files use the flag `-data` for example `./simulation -data=../data`
To use the alternative design use the flag `-alt` for example `./simulation -alt` or `./simulation -alt=true`

## Littles Law
Littles Law tells us that the average number of items(L) in a stationary system is equal to the average arrival rate (λ) multiplied by the average time (W) that a item spends in the system.

Algebraically: L = λW

## Applying Littles Law to our system
For our system we want to model the item as a product and want to maximize thoughput which means minimize how long it spends within the system (min W). To do this we can either decrease the number of items in a the stationary system (L) or increase the arrival rate.

It was noted that products 2 and 3 were being produced significantlly slower than product 1 so, in order to optimize the entire system, the focus was put into optimizing production of products 2 and 3. The assumution was made that simply changing the number of items in the stationary system was not possible because this would require changing the buffer sizes given in the project spec. Therefor to decrease the time spent in the system the arrival rate for the components into ws2 and ws3 was increased, this was done by favouring ws3, ws2, and ws1 in that order when distributing shared components.

The following table shows the average time spent for each product in the system:

| Product    | Original | Alternative |
|------------|----------|-------------|
| Product 1  | 10.37945510626m | 12.208296922943333m |
| Product 2  | 144.55721668924272m | 117.59869925506202m |
| Product 3  | 137.92389689436362m | 117.03486057925582m |
| Sum        | 292.86056868986634m | 246.84185675726118m |
