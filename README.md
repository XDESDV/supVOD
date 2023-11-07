# supVOD
Online video streaming at SupDeVinci

# MODELS
USERS
  email
  first_name
  last_name
  password
  
MOVIES
  Title
  description
  duration
  kinds[]
  
KINDS
  name
  
HISTORICS (duration)
  USER
  MOVIE
  duration

# Ressources
/movies
   Find
   GetbyId
   Create
   Update

/User
   Create
   Update

/kinds
   Find
   Create
   Update

/historics
  Find
  Get
  Create
  update
