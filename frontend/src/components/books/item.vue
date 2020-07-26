<template>
  <div>
    <router-link :to="'/api/books/'+book.id">
      {{ book.title }}
    </router-link>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  props: ['id'],
  data: () => ({
    book: ''
  }),
  mounted() {
    axios
      .get(`http://localhost:8000/api/books/${this.$props.id}`)
      .then(res => (this.book = res.data))
  },
}
</script>
