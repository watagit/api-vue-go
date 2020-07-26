<template>
  <div>
    {{ book.title }}
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data: () => ({
    book: ''
  }),
  mounted() {
    axios
      .get(`http://localhost:8000/api/books/${this.$route.params.id}`)
      .then(res => (this.book = res.data))
  }
}
</script>
