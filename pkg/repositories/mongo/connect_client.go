// Package mongo proporciona utilidades de conexión a MongoDB y configuración de infraestructura.
// Este paquete contiene adaptadores para conectar a MongoDB y sirve como la
// capa de infraestructura en la arquitectura hexagonal.
package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectClient establece una conexión a una instancia de MongoDB usando la URI proporcionada.
// Establece un timeout de 30 segundos para el proceso de conexión y verifica la conexión
// haciendo ping al servidor. Retorna el cliente de MongoDB si es exitoso, o un error si
// la conexión no puede ser establecida.
//
// Parámetros:
//   - dbURI: La URI de conexión a MongoDB (ej., "mongodb://localhost:27017")
//
// Retorna:
//   - client: Una instancia del cliente de MongoDB conectado
//   - err: Un error si la conexión falla o si la verificación de ping falla
func ConnectClient(dbURI string) (client *mongo.Client, err error) {
	// Establece un timeout para permitir que el proceso de conexión se cancele si toma demasiado tiempo.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Conecta al servidor de MongoDB
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		return nil, err
	}

	// Llama al método Ping para verificar que la conexión se ha establecido exitosamente.
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
