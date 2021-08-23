<template>
    <div class='container'>
        test123
    </div>
</template>

<script>
import {testT} from '@/api/test';

export default {
    name: 'test',
    data() {
        return {}
    },
    props: {},
    computed: {},
    created() {
        testT()
    },
    mounted() {},
    activated() {},
    methods: {},
    watch: {},
    components: {},
}
</script>

<style lang='scss' scoped>
// @import '@/style/mixin';

</style>