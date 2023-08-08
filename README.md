# QuizAppApi

This is a simple **REST API** designed for
quiz application made during the lectures in the [PDP University](https://university.pdp.uz).

#### **BY** [Islmobay](https://t.me/+pXNYOnXyHTEwZTUy)

## API Documentation

## Authorization

### Login
Login as an admin

- **URL**: `/api/login`
- **Method**: `POST`
- **Request body**:
```json
{
  "user": "user role",
  "password": "user password"
}
```
- **Response on success**:
```json
{
  "message": "JWT token"
}
```

## Subject API

### Get all subjects

Get a list of all subjects available

- **URL**: `/api/subjects`
- **Method**: `GET`

- **Response**:
```json
[
  {
    "id": "1",
    "name": "Math",
    "iconPath": "https://example.com/math.png",
    "color": "fff333"
  }
]
```

### Get one subject

Get a single subject by providing its ID.

- **URL**: `/api/subjects/{subjectID}`
- **Method**: `GET`
- **Response**:

```json
{
  "id": 1,
  "name": "Math",
  "iconPath": "https://example.com/math.png",
  "color": "fff333",
  "questions": [
    {
      "sid": 1,
      "qid": 1,
      "text": "3 x 4 = ?",
      "answer1": "10",
      "answer2": "12",
      "answer3": "14",
      "answer4": "9",
      "correctAnswer": "12",
      "level": 1
    },
    {
      "sid": 1,
      "qid": 2,
      "text": "3 x 5 = ?",
      "answer1": "12",
      "answer2": "16",
      "answer3": "15",
      "answer4": "25",
      "correctAnswer": "15",
      "level": 2
    }
  ]
}
```
- **Returned values**:
    - `sid` - The ID of the subject to which the question is assigned
    - `qid` - The ID of the question
    - `level` - The level of difficulty of the question

### Add subject

Add subject to the database

- **URL**: `/api/admin/subject`
- **Method**: `POST`
- **Headers**:
  - **Authorization**: `Bearer {JWT token}`
- **Request body**:

```json
{
  "name": "New Subject's Name",
  "color": "new color for the subject",
  "iconPath": "subject.png",
  "questions": [
    {
      "text": "question text",
      "a1": "answer 1",
      "a2": "answer 2",
      "a3": "answer 3",
      "a4": "answer 4",
      "ca": "correct answer",
      "level": "level of difficulty"
    }
  ]
}
```
- **Response on success**:
```json
{
  "message": "id of a new subject"
}
```

### Update subject

Update the existing subject to the new one

- **URL**: `/api/admin/subject`
- **Method**: `PUT`
- **Headers**:
- **Authorization**: `Bearer {JWT token}`

- **Request body**:

```json
{
  "id": "id of the subject",
  "name": "new name for the subject",
  "iconPath": "new icon path for the subject",
  "color": "new color for the color"
}
```
- **Response on success**:

```json
{
  "message": true
}
```

### Delete subject

Delete the subject by ID

- **URL**: `/api/admin/subject`
- **Method**: `DELETE`
- **Headers**:
  - **Authorization**: `Bearer {JWT token}`

- **Request body**:

```json
{
  "id": "id of the subject"
}
```
- **Response on success**:

```json
{
  "message": true
}
```

## Questions API

### Add question
Add new question to the subject

- **URL**: `/api/admin/question`
- **Method**: `POST`
- **Headers**:
  - **Authorization**: `Bearer {JWT token}`

- **Request body**:

```json
{
  "sid": "Subject ID",
  "text": "text of the question",
  "a1": "answer 1",
  "a2": "answer 2",
  "a3": "answer 3",
  "a4": "answer 4",
  "ca": "correct answer",
  "level": "level of difficulty"
}
```

- **Response on success**:
```json
{
  "message": true
}
```

### Update question

Update existing question by ID
- **URL**: `/api/admin/question`
- **Method**: `PUT`
- **Headers**:
  - **Authorization**: `Bearer {JWT token}`

- **Request body**:

```json
{
  "sid": "Subject ID",
  "qid": "Question id",
  "text": "text of the question",
  "a1": "answer 1",
  "a2": "answer 2",
  "a3": "answer 3",
  "a4": "answer 4",
  "ca": "correct answer",
  "level": "level of difficulty"
}
```

- **Response on success**:
```json
{
  "message": true
}
```


### Delete question

Delete existing question by ID
- **URL**: `/api/admin/question`
- **Method**: `DELETE`
- **Headers**:
  - **Authorization**: `Bearer {JWT token}`

- **Request body**:

```json
{
  "sid": "Subject ID",
  "qid": "Question id"
}
```

- **Response on success**:
```json
{
  "message": true
}
```