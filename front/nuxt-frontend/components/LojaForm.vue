<template>
  <form @submit.prevent="salvar">
    <input v-model="form.nome" placeholder="Nome" required />
    <input v-model="form.numero_loja" placeholder="Número" required />
    <button type="submit">Salvar</button>
    <button @click="$emit('fechar')">Cancelar</button>
  </form>
</template>

<script setup>
const props = defineProps(['estabelecimentoId'])
const emit = defineEmits(['salvar', 'fechar'])
const { $api } = useNuxtApp()
const form = reactive({
  nome: '', numero_loja: '', estabelecimento_id: props.estabelecimentoId
})

const salvar = async () => {
  await $api('/lojas', { method: 'POST', body: form })
  emit('salvar')
  emit('fechar')
}
</script>
