import 'package:flutter/material.dart';
import 'package:frontend/screens/login_screen.dart';
import 'package:provider/provider.dart';
import '../../providers/auth_provider.dart';
import '../../providers/ticket_provider.dart';
import '../../widgets/ticket_card.dart';

class AssignedTicketsScreen extends StatefulWidget {
  const AssignedTicketsScreen({super.key});

  @override
  State<AssignedTicketsScreen> createState() => _AssignedTicketsScreenState();
}

class _AssignedTicketsScreenState extends State<AssignedTicketsScreen> {
  @override
  void initState() {
    super.initState();
    WidgetsBinding.instance.addPostFrameCallback((_) {
      context.read<TicketProvider>().loadAssignedTickets();
    });
  }

  Future<void> _updateStatus(String ticketId, String status) async {
    await context.read<TicketProvider>().updateStatus(ticketId, status);
    if (!mounted) return;
    context.read<TicketProvider>().loadAssignedTickets();
  }

  @override
  Widget build(BuildContext context) {
    final tickets = context.watch<TicketProvider>();
    final auth = context.read<AuthProvider>();

    return Scaffold(
      appBar: AppBar(
        title: const Text('Tickets Assignados'),
        actions: [
          IconButton(
            icon: const Icon(Icons.logout),
            onPressed: () async {
              await auth.logout();
              if (!mounted) return;
              Navigator.pushReplacement(
                context,
                MaterialPageRoute(builder: (_) => const LoginScreen()),
              );
            },
          ),
        ],
      ),
      body: tickets.isLoading
          ? const Center(child: CircularProgressIndicator())
          : tickets.tickets.isEmpty
          ? const Center(child: Text('Nenhum ticket assignado'))
          : ListView.builder(
              itemCount: tickets.tickets.length,
              itemBuilder: (context, index) {
                final ticket = tickets.tickets[index];
                return TicketCard(
                  ticket: ticket,
                  onTap: () => showModalBottomSheet(
                    context: context,
                    builder: (_) => Padding(
                      padding: const EdgeInsets.all(24),
                      child: Column(
                        mainAxisSize: MainAxisSize.min,
                        children: [
                          Text(
                            ticket.title,
                            style: const TextStyle(
                              fontSize: 18,
                              fontWeight: FontWeight.bold,
                            ),
                          ),
                          const SizedBox(height: 8),
                          Text(ticket.description),
                          const SizedBox(height: 24),
                          SizedBox(
                            width: double.infinity,
                            child: ElevatedButton(
                              onPressed: () {
                                Navigator.pop(context);
                                _updateStatus(ticket.id, 'DONE');
                              },
                              child: const Text('Marcar como Concluído'),
                            ),
                          ),
                        ],
                      ),
                    ),
                  ),
                );
              },
            ),
    );
  }
}
