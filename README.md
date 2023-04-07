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
    
-> Definir atributos que tendrá un mensaje, tales como; TipoDeCarta, CantidadDeCartas,

v1.0 

 Trigger de la funcion mediante mensaje de pubsub <----> Adaptadores de entrada (Pubsub) que recibirá el mensaje y lo enviará a un coordinador <---> Coordinar que recibe el mensaje y lo manda al adaptador 
 que busca la información  <---> Adaptador de busqueda de información que busca y devuelve lo solicitado

 v1.1 

 Trigger de la funcion mediante mensaje de pubsub el cual tendrá el mensaje y atributos según la carta solicitada 
 <----> 
 Adaptador de busqueda de información que recibe la solicitud y devuelve la cantidad de cartas solicitada

 La nueva versión se hará en otro proyecto nuevo, esto se dejará como una maqueta para futuros proyectos.