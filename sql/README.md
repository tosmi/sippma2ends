# Hinweise zum Datenexport

* sippma_er.png ist das verwendete ER Modell
  * wir haben nur 
    * Patients (enthaelt Patienten und Eltern)
    * Consultations und
    * Relationships exportiert
  * Relationships enthaelt die Verknuepfung von Patienten zu deren Eltern
* Leider sind in der Patienten Tabelle die Eltern teilweise doppelt vorhanden
  (nur bei Geschwisterkindern)
  * D.h. es gibt mehrere Eintraege mit der gleichen Versicherungsnummer (SSN)
  * Ein Moeglichkeit waere die doppelten Eintraege einfach zu ignorieren und die 
    Patienten, die danach keinen Eltern mehr zugwiesen haben, manuell in Latido
    zu beheben. 