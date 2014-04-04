# Commands

## TakeOff

Sets the internal `fly` state to `true`.

## Land

Sets the internal `fly` state to `false`.

## Up(speed float64) 

Makes the drone gain altitude. 
speed can be a value from `0.0` to `1.0`.

##### Params

- **speed** - **float64** - The speed at which the drone moves

## Down(speed float64) 

Makes the drone reduce altitude. 
speed can be a value from `0.0` to `1.0`.

##### Params

- **speed** - **float64** - The speed at which the drone moves

## Left(speed float64) 

Causes the drone to bank to the left, controls the roll, which is 
a horizontal movement using the camera as a reference point.
speed can be a value from `0.0` to `1.0`.

##### Params

- **speed** - **float64** - The speed at which the drone moves

## Right(speed float64) 

Causes the drone to bank to the right, controls the roll, which is 
a horizontal movement using the camera as a reference point.
speed can be a value from `0.0` to `1.0`.

##### Params

- **speed** - **float64** - The speed at which the drone moves

## Forward(speed float64)

Causes the drone go forward, controls the pitch.
speed can be a value from `0.0` to `1.0`.

##### Params

- **speed** - **float64** - The speed at which the drone moves


## Backward(speed float64)

Causes the drone go backward, controls the pitch.
speed can be a value from `0.0` to `1.0`.

##### Params

- **speed** - **float64** - The speed at which the drone moves

## Clockwise(speed float64)

Causes the drone to spin.
speed can be a value from `0.0` to `1.0`.

##### Params

- **speed** - **float64** - The speed at which the drone moves

## CounterClockwise(speed float64)

Causes the drone to spin.
speed can be a value from `0.0` to `1.0`.

##### Params

- **speed** - **float64** - The speed at which the drone moves

## Hover

Causes the drone to hover in place.
