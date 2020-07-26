<template>
  <div id="app">
    <h1>books</h1>
    <BookItem
      v-for="book in books"
      :key="book.id"
      :id="book.id"
    />
  </div>
</template>

<script>
import axios from 'axios';

import BookItem from '../../components/books/item'

export default {
  components: {
    BookItem
  },
  data: () => ({
    books: '',
  }),
  mounted() {
    axios
      .get("http://localhost:8000/api/books")
      .then(res => (this.books = res.data))
  },
}
</script>
