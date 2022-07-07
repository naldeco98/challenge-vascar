# challenge-vascar
Coding challange for vascar solutions

### An application has a social posts and comments feature. Both, posts and comments can be reported by users. The report needs to store,

·         A reason why the user reported it (free text)

·         The user who reported it (a user id, for the purpose of this practice program just a random integer)

·         The date when the report was made

### Posts and comments are very simple, they just store

·         A unique identifier (integer)

·         A text field (free text limited to 500 characters)

·         The post/comment creation date

### Write a simple application that reads comments from a sqlite database and writes these reports back to it through a REST API. You will only need 2 endpoints, namely the one for reporting comments and the one for reporting posts. You will also need to write simple end-to-end tests using the native Go testing utilities.