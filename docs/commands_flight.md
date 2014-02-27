# Commands

## TakeOff

Sets the internal `fly` state to `true`.


##### Returns 

`nil`

## Land

Sets the internal `fly` state to `false`.

##### Returns 

`nil`

## Up(speed) 

Makes the drone gain altitude. 
speed can be a value from `0` to `1`.

##### Params

- **speed** - The speed for which to return data

##### Returns 

`integer`

## Down(speed) 

Makes the drone reduce altitude. 
speed can be a value from `0` to `1`.

##### Params

- **speed** - The speed for which to return data

##### Returns 

`integer`

## Left(speed) 

Causes the drone to bank to the left, controls the roll, which is 
a horizontal movement using the camera as a reference point.
speed can be a value from `0` to `1`.

##### Params

- **speed** - The speed for which to return data

##### Returns 

`integer`

## Right(speed) 

Causes the drone to bank to the right, controls the roll, which is 
a horizontal movement using the camera as a reference point.
speed can be a value from `0` to `1`.

##### Params

- **speed** - The speed for which to return data

##### Returns 

`integer`

## Forward

Causes the drone to go forward.

##### Returns 

`nil`

## Backward

Causes the drone to go forward.

##### Returns 

`nil`

## Clockwise(speed)

Causes the drone to spin.
speed can be a value from `0` to `1`.

##### Params

- **speed** - The speed for which to return data

##### Returns 

`integer`

## CounterClockwise(speed)

Causes the drone to spin.
speed can be a value from `0` to `1`.

##### Params

- **speed** - The speed for which to return data

##### Returns 

`integer`

## Hover

Causes the drone to go hover.

##### Returns 

`nil`

