package model

import (
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

// Embedding Defaults
const (
	EmbeddingModelTag = "serve"
	EmbeddingOp       = "StatefulPartitionedCall"
)

// NewEmbeddings returns a pre-trained model for text embeddings
func NewEmbeddings(path string, opts ...BertOption) (Bert, error) {
	vocabPath := (path + "/" + DefaultVocabFile) // TODO os.Join
	m, err := tf.LoadSavedModel(path, []string{EmbeddingModelTag}, nil)
	if err != nil {
		return Bert{}, err
	}
	return NewBert(m, vocabPath, opts...)
}
