<template>
  <div class="p-6 max-w-2xl mx-auto">
    <h1 class="text-xl font-bold mb-4">Editar Loja</h1>

    <form v-if="form" @submit.prevent="atualizarLoja" class="space-y-4">
      <div class="grid grid-cols-2 gap-4">
        <div>
          <label class="block mb-1 font-semibold">Número da Loja</label>
          <input v-model="form.numero_loja" class="w-full border p-2 rounded" required />
        </div>

        <div>
          <label class="block mb-1 font-semibold">Nome</label>
          <input v-model="form.nome" class="w-full border p-2 rounded" required />
        </div>

        <div>
          <label class="block mb-1 font-semibold">Razão Social</label>
          <input v-model="form.razao_social" class="w-full border p-2 rounded" required />
        </div>

        <div>
          <label class="block mb-1 font-semibold">Endereço</label>
          <input v-model="form.endereco" class="w-full border p-2 rounded" required />
        </div>

        <div>
          <label class="block mb-1 font-semibold">Número</label>
          <input v-model="form.numero" class="w-full border p-2 rounded" required />
        </div>

        <div>
          <label class="block mb-1 font-semibold">Cidade</label>
          <input v-model="form.cidade" class="w-full border p-2 rounded" required />
        </div>

        <div>
          <label class="block mb-1 font-semibold">Estado</label>
          <input v-model="form.estado" class="w-full border p-2 rounded" required />
        </div>

        <div>
          <label class="block mb-1 font-semibold">CEP</label>
          <input v-model="form.cep" class="w-full border p-2 rounded" required />
        </div>

        <div class="col-span-2">
          <label class="block mb-1 font-semibold">Estabelecimento</label>
          <select v-model="form.estabelecimento_id" class="w-full border p-2 rounded" required>
            <option value="" disabled>Selecione um estabelecimento</option>
            <option v-for="estabelecimento in estabelecimentos" :key="estabelecimento.id" :value="estabelecimento.id">
              {{ estabelecimento.nome }}
            </option>
          </select>
        </div>
      </div>

      <button
        type="submit"
        class="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700"
      >
        Atualizar
      </button>
    </form>

    <p v-else class="text-gray-500">Carregando dados da loja...</p>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter, useNuxtApp } from '#app'

// Tipagens
interface Loja {
  id?: number
  numero_loja: string
  nome: string
  razao_social: string
  endereco: string
  cidade: string
  estado: string
  cep: string
  numero: string
  estabelecimento_id: number | null
}

interface Estabelecimento {
  id: number
  nome: string
}

const route = useRoute()
const router = useRouter()
const { $api } = useNuxtApp()

const form = ref<Loja | null>(null)
const estabelecimentos = ref<Estabelecimento[]>([])

const carregarEstabelecimentos = async () => {
  try {
    const response = await $api<{ data: Estabelecimento[] }>('/estabelecimentos')
    estabelecimentos.value = response.data
  } catch (error) {
    alert('Erro ao carregar estabelecimentos')
    console.error(error)
  }
}

const carregarLoja = async () => {
  const id = route.params.id
  try {
    const response = await $api<{ data: Loja }>(`/lojas/${id}`)
    form.value = response.data
  } catch (error) {
    alert('Erro ao carregar loja')
    console.error(error)
  }
}

const atualizarLoja = async () => {
  if (!form.value) return

  try {
    await $api(`/lojas/${form.value.id}`, {
      method: 'PUT',
      body: form.value
    })
    alert('Loja atualizada com sucesso!')
    await router.push('/estabelecimentos')
  } catch (error: any) {
    alert('Erro ao atualizar loja')
    console.error('Erro detalhado:', error?.response?._data || error)
  }
}

onMounted(async () => {
  await carregarEstabelecimentos()
  await carregarLoja()
})
</script>
