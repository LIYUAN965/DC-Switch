<template>
    <li>props: {{ todo.text }}</li>
    <li>data: {{ todo1.text }}</li>
</template>

<script>
export default {
    props: ['todo'],
    data() {
        console.log(this.todo);
        return {
            todo1: this.todo  // props中定义了，data中不用重复返回
        }
    },
}
</script>