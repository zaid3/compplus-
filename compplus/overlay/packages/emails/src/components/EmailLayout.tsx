import {
  Body,
  Container,
  Head,
  Html,
  Link,
  Section,
  Text,
} from 'react-email';
import * as React from 'react';
import { Logo } from './Logo';
import { ProboLogo } from './ProboLogo';

interface EmailLayoutProps {
  subject: string;
  children: React.ReactNode;
}

export const EmailLayout = ({
  subject,
  children,
}: EmailLayoutProps) => {
  const brandedSubject = subject.replace(/^Probo\b/, 'ISOPilot').replace(/^Comp Plus\+\b/, 'ISOPilot').replace(/^ISOpilot\b/, 'ISOPilot');

  return (
    <Html lang="en">
      <Head>
        <meta name="x-apple-disable-message-reformatting" />
        <meta httpEquiv="X-UA-Compatible" content="IE=edge" />
        <title>{brandedSubject}</title>
      </Head>
      <Body style={main}>
        <Container style={container}>
          <Section style={content}>
            <Section style={logoSection}>
              <Link href="{{.SenderCompanyWebsiteURL}}">
                <Logo />
              </Link>
            </Section>

            <Text style={text}>Hi {'{{.RecipientFullName}}'},</Text>

            {children}
          </Section>

          <Section style={footerSection}>
            <Text style={footerAddress}>
              {"{{.SenderCompanyHeadquarterAddress}}"}
            </Text>
            <Text style={footerAddress}>
              <span style={{verticalAlign: "middle"}}>Your compliance co-pilot · </span>
              <Link style={{display: "inline-block", height: "16px", verticalAlign: "middle"}} href="{{.SenderCompanyWebsiteURL}}">
                <ProboLogo />
              </Link>
            </Text>
          </Section>
        </Container>
      </Body>
    </Html>
  );
};

export default EmailLayout;

const main: React.CSSProperties = {
  margin: '0',
  padding: '0',
  fontFamily: "Inter, -apple-system, BlinkMacSystemFont, 'Segoe UI', Arial, sans-serif",
  backgroundColor: '#F6F8FB',
  WebkitTextSizeAdjust: '100%',
};

const container: React.CSSProperties = {
  maxWidth: '600px',
  width: '100%',
  backgroundColor: '#FFFFFF',
  borderRadius: '12px',
  boxShadow: '0 1px 2px rgba(11,18,32,.06)',
  margin: '40px auto',
  border: '1px solid #E4E9F2',
};

export const headerSection: React.CSSProperties = {
  padding: '40px 40px 30px 40px',
  textAlign: 'center',
  backgroundColor: '#0B1220',
  borderRadius: '12px 12px 0 0',
};

export const h1: React.CSSProperties = {
  margin: '0',
  color: '#FFFFFF',
  fontSize: '24px',
  fontWeight: '700',
  lineHeight: '1.3',
};

const content: React.CSSProperties = {
  padding: '40px',
};

const text: React.CSSProperties = {
  margin: '0 0 20px 0',
  color: '#0B1220',
  fontSize: '16px',
  lineHeight: '24px',
};

export const buttonContainer: React.CSSProperties = {
  padding: '10px 0 30px 0',
};

export const button: React.CSSProperties = {
  display: 'inline-block',
  padding: '14px 28px',
  backgroundColor: '#2F6BFF',
  color: '#FFFFFF',
  textDecoration: 'none',
  borderRadius: '8px',
  fontWeight: '600',
  fontSize: '16px',
  lineHeight: '20px',
};

export const footerText: React.CSSProperties = {
  margin: '0',
  color: '#64748B',
  fontSize: '14px',
  lineHeight: '20px',
};

const footerSection: React.CSSProperties = {
  padding: '30px 40px',
  borderTop: '1px solid #E4E9F2',
};

const footerAddress: React.CSSProperties = {
  margin: '10px 0 0 0',
  color: '#64748B',
  fontSize: '12px',
  lineHeight: '18px',
};

const logoSection: React.CSSProperties = {
  marginBottom: '30px',
};

export const bodyText: React.CSSProperties = {
  margin: '0 0 30px 0',
  color: '#0B1220',
  fontSize: '16px',
  lineHeight: '24px',
};
