/*Package weather provides the information
about the weather conditions of the country 
based on the location and the current location.
*/
package weather

var (
    //CurrentCondition stores the current weather condition.
	CurrentCondition string
    //CurrentLocation stores the country's current location.
	CurrentLocation  string
)
/*Forecast function uses  the data stored in the 
CurrentCondition and CurrentLocation to forecast the 
weather conditions of the country.
*/
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
