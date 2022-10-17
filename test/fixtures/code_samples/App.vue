<!-- Convert kilometers to miles and back. -->

<style scoped>
  label {
    display: block;
    margin: 1em 0;
    width: 20em;
    font-size: 0.8em;
  }
  input {
    width: 10em
  }
  div {
    width: 7em;
    text-align: center;
  }
</style>

<script setup>
import { ref } from 'vue'

const kms = ref(1.609344)
const miles = ref(1)

function setKms(e, v = +e.target.value) {
  kms.value = v
  miles.value = v / 1.609344
}

function setMiles(e, v = +e.target.value) {
  miles.value = v
  kms.value = v * 1.609344
}
</script>

<!--
I would walk 500 miles,
and I would walk 804.672 kilometers...
-->

<template>
  <label>
    <input type="number" :value="kms" @change="setKms">
    Kilometers
  </label>
  <div>⇅</div>
  <label>
    <input type="number" :value="miles" @change="setMiles">
    Miles
  </label>
</template>
