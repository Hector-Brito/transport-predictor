CREATE TABLE IF NOT EXISTS transport_log (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
/* SQLITE maneja las fechas como integer o text*/
  departure_date DATETIME NOT NULL,
  arrival_date DATETIME NOT NULL,
  latency INTEGER,
  vehicle_id INTEGER NOT NULL,
  weather_id INTEGER NOT NULL,
  day_of_week VARCHAR(30) NOT NULL CHECK (day_of_week IN ("Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday")),
  is_fortnight INTEGER NOT NULL DEFAULT 0 CHECK (is_fortnight IN (0, 1)),
  observations TEXT ,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_vehicle_transport_log FOREIGN KEY (vehicle_id) REFERENCES vehicle(id) ON DELETE RESTRICT,
  CONSTRAINT fk_weather_transport_log FOREIGN KEY (weather_id) REFERENCES weather(id) ON DELETE RESTRICT,
  CONSTRAINT check_date CHECK (arrival_date > departure_date)
)
