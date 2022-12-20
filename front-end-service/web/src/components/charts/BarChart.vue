<template>
    <div>
        <v-chart :options="option"/>
    </div>
</template>

<script>
import Echarts from 'vue-chartjs';
import 'echarts-gl';
import 'axios';

export default {
    components: {
        'v-chart': Echarts
    }
};
</script>

<style>
</style>