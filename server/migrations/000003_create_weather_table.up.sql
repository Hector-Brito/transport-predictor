CREATE TABLE IF NOT EXISTS weather (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name VARCHAR(50) NOT NULL DEFAULT "Sunny" CHECK (name IN ("Sunny", "Foggy", "Drizzle", "Storm")),
  ground_status VARCHAR(50) NOT NULL DEFAULT "Dry" CHECK ( ground_status IN ("Dry", "Wet", "Puddles/Mud")),
  visibility VARCHAR(50) NOT NULL DEFAULT "Clear" CHECK (visibility IN ("Clear", "Foggy")),
  intensity INTEGER NOT NULL DEFAULT 0 CHECK (intensity >= 0 AND intensity <= 5),
  temperature INTEGER NOT NULL CHECK (temperature >= -100 AND temperature <= 100),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
