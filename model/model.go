package model


type Admin struct {
    ID        string `json:"id"`
    FirstName string `json:"firstname"`
    LastName  string `json:"lastname"`
    Email     string `json:"email"`
    Password  string `json:"password"`
}


type Customer struct {
    ID        string `json:"id"`
    FirstName string `json:"firstname"`
    LastName  string `json:"lastname"`
    Email     string `json:"email"`
    Password  string `json:"password"`
}


type HairSalon struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Address     string `json:"address"`
    Description string `json:"description"`
    IsAccepted  bool   `json:"is_accepted"`
}


type Hairdresser struct {
    ID          string `json:"id"`
    FirstName   string `json:"firstname"`
    LastName    string `json:"lastname"`
    Email       string `json:"email"`
    Password    string `json:"password"`
    StartTime   string `json:"start_time"` // Format HH:MM:SS
    EndTime     string `json:"end_time"`   // Format HH:MM:SS
    HairSalonID string `json:"hair_salon_id"`
}


type Category struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}


type Service struct {
    ID           string `json:"id"`
    Name         string `json:"name"`
    Price        int    `json:"price"`
    Duration     int    `json:"duration"`     // Durée en minutes
    CategoryID   string `json:"category_id"`
    HairSalonID  string `json:"hair_salon_id"`
}


type Slot struct {
    ID            string `json:"id"`
    StartTime     string `json:"start_time"` // Format "YYYY-MM-DD HH:MM:SS"
    EndTime       string `json:"end_time"`   // Format "YYYY-MM-DD HH:MM:SS"
    IsBooked      bool   `json:"is_booked"`
    HairdresserID string `json:"hairdresser_id"`
}


type Booking struct {
    ID          string `json:"id"`
    IsConfirmed bool   `json:"is_confirmed"`
    CustomerID  string `json:"customer_id"`
    ServiceID   string `json:"service_id"`
    SlotID      string `json:"slot_id"`
}
