<template>
  <div id="app">
    <h1>books</h1>
    <table border="1">
      <tr>
        <th>ID</th>
        <th>Isbn</th>
        <th>Title</th>
        <th>Author</th>
      </tr>
      <tr v-for="book in books" :key="book.id">
        <td>{{ book.id }}</td>
        <td>{{ book.isbn }}</td>
        <td>{{ book.title }}</td>
        <td>{{ book.author.firstname }} {{ book.author.lastname }}</td>
      </tr>
    </table>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      books: '',
    }
  },
  mounted() {
    axios
      .get("http://localhost:8000/api/books")
      .then(res => (this.books = res.data))
  },
}
</script>
