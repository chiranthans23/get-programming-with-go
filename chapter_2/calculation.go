//Malacandra is another name for Mars in The Space Trilogy by C. S. Lewis.
//Write a program to determine how fast a ship would need to travel (in km/h) in order to reach Malacan- dra in 28 days.
// Assume a distance of 56,000,000 km.
package chapter_2

import "fmt"

//CalSpeedToMars defines parameters for mars and calculates the speed needed
func CalSpeedToMars() {
	var days = 28
	const dist = 56000000
	fmt.Printf("Speed is (distance/time) and the speed required to travel '%v' distance in '%v' days is '%v' km/day\n", dist, days, (dist / days))

}
