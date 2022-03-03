// Not an executable file, but a package for organizing functions.
package shows

// Import libraries.
import "math/rand"

// Add a collection of show names, from which we will randomly select when the
// client requests a show recommendation.
var showNames = [...]string {
  "Peaky Blinders",
  "Daredevil",
  "The Witcher",
  "The Haunting of Hill House",
  "Castlevania",
  "BoJack Horseman",
  "New Girl",
  "Midnight Mass",
}

// Note that functions which begin with a capital letter are understood
// to be "public", so that they can be invoked elsewhere. Functions that
// begin with a lowercase letter are private and will raise an error if
// invoked outside of this package.
func SelectRandom() string {
  // Fetch a random index for our `showNames` array.
  randomIndex := rand.Intn(len(showNames))

  // Select the random item from our collection of show names.
  name := showNames[randomIndex]

  // Return the resulting show name.
  return name
}
