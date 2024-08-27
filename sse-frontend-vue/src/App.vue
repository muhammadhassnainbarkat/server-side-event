<script setup>
import {onUnmounted, ref} from "vue";

const error = ref(null)
const messages = ref([]);
const eventSource = ref(null);
const clientId = "1";

const subscribe = () => {

  if(eventSource.value) {return}

  eventSource.value = new EventSource("http://localhost:8080/event/subscribe?clientId=" + clientId);
  eventSource.value.onmessage = (event) => {
    messages.value.push(event.data)
  }

  eventSource.value.onerror = (event) => {
    this.error.value = event.error
  }

}

const disconnect = () => {
  if(eventSource.value){
    fetch('http://localhost:8080/event/unsubscribe?clientId=' + clientId)
        .then(() => {
          eventSource.value.close();
          messages.value = [];
        })
  }
}

onUnmounted(()=> {
  disconnect();
})

</script>

<template>
  <main>
    <div class="container">
      <h1> Server side event</h1>

      <p v-if="error">{{error}}</p>
      <div v-else>
        <button @click="subscribe">Subscribe</button>
        <button @click="disconnect">Unsubscribe</button>
        <div v-for="(message, index) in messages" :key="index">
          <p>{{ message }}</p>
        </div>

      </div>
    </div>
  </main>
</template>

<style scoped>
</style>
