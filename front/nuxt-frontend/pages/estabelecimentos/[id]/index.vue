<template>
  <div class="p-6 max-w-2xl mx-auto">
    <h1 class="text-xl font-bold mb-4">Editar Estabelecimento</h1>

    <EstabelecimentoForm
      v-if="estabelecimento"
      :estabelecimento="estabelecimento"
      @submit="salvarEstabelecimento"
    />

    <p v-else class="text-gray-500">Carregando dados...</p>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter, useNuxtApp } from '#app'
import EstabelecimentoForm from '@/components/EstabelecimentoForm.vue'

interface Estabelecimento {
  id?: number
  numero_estabelecimento: string
  nome: string
  razao_social: string
  endereco: string
  cidade: string
  estado: string
  cep: string
  numero: string
}

const route = useRoute()
const router = useRouter()
const { $api } = useNuxtApp()

const estabelecimento = ref<Estabelecimento | null>(null)

const carregarEstabelecimento = async () => {
  const id = route.params.id

  const estado = window.history.state?.estabelecimento
  if (estado) {
    estabelecimento.value = estado
    return
  }

  try {
    const response = await $api<{ data: Estabelecimento }>(`/estabelecimentos/${id}`)
    estabelecimento.value = response.data
  } catch (error) {
    alert('Erro ao carregar estabelecimento')
    console.error(error)
  }
}

const salvarEstabelecimento = async (form: Estabelecimento) => {
  try {
    await $api(`/estabelecimentos/${form.id}`, {
      method: 'PUT',
      body: form
    })
    alert('Estabelecimento atualizado com sucesso!')
    await router.push('/estabelecimentos')
  } catch (error: any) {
    alert('Erro ao salvar estabelecimento')
    console.error('Erro detalhado:', error?.response?._data || error)
  }
}

onMounted(carregarEstabelecimento)
</script>
