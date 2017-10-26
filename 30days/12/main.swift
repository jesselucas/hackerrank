import Foundation

// Class Person
class Person {
    var firstName: String
    var lastName: String
    var id: Int

    // Initializer
    init(firstNameString: String, lastNameString: String, identificationNumber: Int) {
        firstName = firstNameString
        lastName = lastNameString
        id = identificationNumber
    }

    // Print person data
    func printPerson() {
        print("Name: \(lastName), \(firstName)")
        print("ID: \(id)")
    }
} // End of class Person

// Class Student
class Student: Person {
    var testScores: [Int]

    /*	
    *   Initializer
    *   
    *   Parameters:
    *   firstName - A string denoting the Person's first name.
    *   lastName - A string denoting the Person's last name.
    *   id - An integer denoting the Person's ID number.
    *   scores - An array of integers denoting the Person's test scores.
    */
    // Write your initializer here
    init(firstName: String, lastName: String, identification: Int, scores: [Int]) {
        self.testScores = scores
        super.init(firstNameString: firstName, lastNameString: lastName, identificationNumber: identification)

        self.firstName = firstName
        self.lastName = lastName
        self.id = id
    }

    /*	
    *   Method Name: calculate
    *   Return: A character denoting the grade.
    */
    // Write your method here
    func calculate() -> Character {
    	var grade:Character
        grade = "T"
        
        var sum = 0
        for s in self.testScores {
            sum += s
        }
        let avg = sum/self.testScores.count
        if 90 <= avg && avg <= 100 {
            grade = "O"
        } else if 80 <= avg && avg < 90 {
            grade = "E"
        } else if 70 <= avg && avg < 80 {
            grade = "A"
        } else if 55 <= avg && avg < 70 {
            grade = "P"
        } else if 40 <= avg && avg < 55 {
            grade = "D"
        }

        return grade
    }
} // End of class Student

if let rl = readLine() {
    let nameAndID = rl.components(separatedBy: " ").map{ String($0) }

    let _ = readLine() // Ignore number of scores input

    if let rl = readLine() {
        let scores = rl.components(separatedBy: " ").map{ Int($0)! }

        let s = Student(firstName: nameAndID[0], lastName: nameAndID[1], identification: Int(nameAndID[2])!, scores: scores)

        s.printPerson()

        print("Grade: \(s.calculate())")
    } else {
         print("need input")
    }

} else {
    print("need input")
}