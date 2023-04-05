# RandomDataGenerator

Esta pieza se comunicará con el core de la pagina web mediante mensajes pubsub, quitando asi
la pesada logica de revisar y hacer multiples consultas a firestore y simplificandola como un adaptador al nucleo principal y agilizandolo. 

Core <---> Pub/sub topic <---> RandomDataGenerator <---> Firestore

todo:
-> Hacer un modelo de como se almacenarán los datos de las cartas en firestore
    -> JSON de ejemplo de cada carta

-> Actualizar la cloud function que almacena cartas de tipo POKEMON añadiendo metodos para cada tipo de carta
    -> Crear y guardar cartas
    --> Actualmente existe la CF en local, falta actualizar el .env para incluir las otras cartas (Para pokemons está listo).
    


