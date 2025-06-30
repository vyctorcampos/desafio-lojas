<template>
  <div class="p-6 max-w-2xl mx-auto">
    <h1 class="text-xl font-bold mb-4">Novo Estabelecimento</h1>
    <EstabelecimentoForm
      :estabelecimento="estabelecimento"
      @submit="criarEstabelecimento"
    />
  </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
import { useRouter, useNuxtApp } from '#app'
import EstabelecimentoForm from '~/components/EstabelecimentoForm.vue'

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

const router = useRouter()
const { $api } = useNuxtApp()

// estado inicial
const estabelecimento = reactive<Estabelecimento>({
  numero_estabelecimento: '',
  nome: '',
  razao_social: '',
  endereco: '',
  cidade: '',
  estado: '',
  cep: '',
  numero: ''
})

// função para enviar os dados
const criarEstabelecimento = async (form: Estabelecimento) => {
  try {
    // remove propriedades reativas (Proxy)
    const payload = JSON.parse(JSON.stringify(form))

    const response = await $api<{ data: Estabelecimento }>('/estabelecimentos', {
      method: 'POST',
      body: payload
    })

    alert('Estabelecimento criado com sucesso!')
    await router.push('/estabelecimentos')
  } catch (error: any) {
    console.error('Erro ao criar estabelecimento:', error?.response?._data || error)
    alert('Erro ao criar estabelecimento')
  }
}
</script>
