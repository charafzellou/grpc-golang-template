package main

import "log"

func checkWinner(srv *server) {
	maxReputation := int32(0)
	maxReputationClient := ""
	for client, reputation := range srv.clients {
		if reputation > maxReputation {
			maxReputation = reputation
			maxReputationClient = client
		}
	}
	if maxReputationClient != "" {
		if srv.clientsFunds[maxReputationClient] > 1000 {
			ascii_reward()
			log.Printf("Client %s has max reputation: %d", maxReputationClient, maxReputation)
			log.Printf("Client %s has won the game with %d,00 €", maxReputationClient, srv.clientsFunds[maxReputationClient])
			log.Printf("Client %s is the winner ! ", maxReputationClient)
			log.Fatal("Game over !")
		}
	}
}

func ascii_reward() {
	log.Printf("")
	log.Printf("  #  ............................................................................................................................................")
	log.Printf("  #  ............................................................................................................................................")
	log.Printf("  #  ██....██..██████..██....██......█████..██████..███████.....████████.██...██.███████.....██.....██.██.███....██.███....██.███████.██████..██.")
	log.Printf("  #  .██..██..██....██.██....██.....██...██.██...██.██.............██....██...██.██..........██.....██.██.████...██.████...██.██......██...██.██.")
	log.Printf("  #  ..████...██....██.██....██.....███████.██████..█████..........██....███████.█████.......██..█..██.██.██.██..██.██.██..██.█████...██████..██.")
	log.Printf("  #  ...██....██....██.██....██.....██...██.██...██.██.............██....██...██.██..........██.███.██.██.██..██.██.██..██.██.██......██...██....")
	log.Printf("  #  ...██.....██████...██████......██...██.██...██.███████........██....██...██.███████......███.███..██.██...████.██...████.███████.██...██.██.")
	log.Printf("  #  ............................................................................................................................................")
	log.Printf("  #  ............................................................................................................................................")
	log.Printf("")
}
